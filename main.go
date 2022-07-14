package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	// "strings"
	"time"
)

var debug_mode bool

func create_github_PR_request() error {
	var pr_arr []interface{}
	res, err := http.Get("https://api.github.com/repos/SarahNerdyBird/Go/issues")
	if err != nil {
		log_error(err.Error())
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log_error(fmt.Sprintf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body))
	}
	if err != nil {
		log_error(err.Error())
	}
	json.Unmarshal(body, &pr_arr)

	fmt.Printf("%s", pr_arr[0])
	return nil
}

// func create_github_PR_request() error {
// 	var body_contents string
// 	var response_body []byte
// 	// TODO: fill in body_contents
// 	body_contents = "{owner: 'SarahNerdyBird', repo: 'Go'}";
// 	body := strings.NewReader(body_contents)

// 	get_request, error := http.NewRequest(
// 		"GET",
// 		"https://api.github.com/repos/SarahNerdyBird/Go/issues",
// 		body,
// 	)
// 	if error != nil {
// 		log_error("Failed to create the HTTP request for the GitHub API")
// 		panic(error.Error())
// 	}

// 	get_request.Header.Set("Accept", "application/vnd.github+json")
// 	// TODO remove token
// 	get_request.Header.Set("Authorization", "token ***")

// 	post_response, error := http.DefaultClient.Do(get_request)
// 	if error != nil {
// 		log_error("Failed to send the HTTP request for the GitHub API")
// 		panic(error.Error())
// 	}

// 	defer post_response.Body.Close()

// 	if post_response.StatusCode < 200 || post_response.StatusCode > 299 {
// 		error_message := fmt.Sprintf(
// 			"Received a non 2xx status code from GitHub. StatusCode: %d. Status: %s",
// 			post_response.StatusCode,
// 			post_response.Status,
// 		)
// 		log_error(error_message)
// 		return fmt.Errorf(error_message)
// 	}

// 	log_info(
// 		fmt.Sprintf(
// 			"GitHub PRs have been retrieved. Received StatusCode: %d and Status: %s.",
// 			post_response.StatusCode,
// 			post_response.Status,
// 		),
// 	)

// 	_, _ = post_response.Body.Read(response_body);

// 	log_info(
// 		fmt.Sprintf(
// 			"Response: %s",
// 			response_body,
// 		),
// 	)
// 	return nil
// }

func log_error(message string) {
	log_message("error: ", message)
}

func log_debug(message string) {
	if debug_mode {
		log_message("debug: ", message)
	}
}

func log_info(message string) {
	log_message("info:  ", message)
}

// Formats all the log messages and adds a timestamp to them.
func log_message(log_level string, message string) {
	// Gets the current time in UTC and rounds to the nearest second.
	timestamp := time.Now().UTC().Round(time.Second).Format("2006-01-02 15:04:05")
	fmt.Printf("%s %s%s\n", timestamp, log_level, message)
}

func main() {
	// Determine the mode in which to run the VOPR Hub
	flag.BoolVar(&debug_mode, "debug", false, "enable debug logging")
	flag.Parse()

	err := create_github_PR_request()
	if err != nil {
		log_error("Failed to get GitHub PRs.")
		return
	}
}

// // TODO change to checkout branch - need a pull command that prefers their changes
// // Note: Every pull request is an issue, but not every issue is a pull request. For this reason, "shared" actions for both features, like manipulating assignees, labels and milestones, are provided within the Issues API.
// // Fetch available branches from GitHub and checkout the correct commit if it exists.
// func checkout_commit(commit string, message_hash []byte) error {
// 	// Ensures commit is all hexadecimal.
// 	commit_valid, error := regexp.MatchString(`^([0-9a-f]){40}$`, commit)
// 	if error != nil {
// 		error_message := fmt.Sprintf("Regex failed to run on the GitHub commit %s", error.Error())
// 		log_error(error_message, message_hash)
// 		return error
// 	} else if !commit_valid {
// 		error = fmt.Errorf("The GitHub commit contained unexpected characters")
// 		log_error(error.Error(), message_hash)
// 		return error
// 	}

// 	// Git commands need to be run with the TigerBeetle directory as their working_directory
// 	fetch_command := exec.Command("git", "fetch", "--all")
// 	fetch_command.Dir = tigerbeetle_directory
// 	error = fetch_command.Run()
// 	if error != nil {
// 		error_message := fmt.Sprintf("Failed to run git fetch: %s", error.Error())
// 		log_error(error_message, message_hash)
// 		return error
// 	}

// 	// Checkout the commit specified in the vopr_message
// 	checkout_command := exec.Command("git", "checkout", commit)
// 	checkout_command.Dir = tigerbeetle_directory
// 	error = checkout_command.Run()
// 	if error != nil {
// 		error_message := fmt.Sprintf("Failed to run git checkout: %s", error.Error())
// 		log_error(error_message, message_hash)
// 		return error
// 	}

// 	// Inspect the git logs.
// 	log_command := exec.Command("git", "log", "-1")
// 	log_command.Dir = tigerbeetle_directory
// 	log_output := make([]byte, 47)
// 	log_output, error = log_command.Output()
// 	if error != nil {
// 		error_message := fmt.Sprintf("Failed to run git log: %s", error.Error())
// 		log_error(error_message, message_hash)
// 		return error
// 	}

// 	// Check the log to determine if the commit has been successfully checked out.
// 	current_commit := string(log_output[0:47])
// 	checkout_successful, error := regexp.MatchString("^commit "+commit, current_commit)
// 	if error != nil {
// 		error_message := fmt.Sprintf(
// 			"Regular expression failure while checking the git logs: %s",
// 			error.Error(),
// 		)
// 		log_error(error_message, message_hash)
// 		return error
// 	}

// 	if !checkout_successful {
// 		error = fmt.Errorf("Checkout failed")
// 		return error
// 	}

// 	return nil
// }

// // TO query the PRs
// // Get PRs, get labels, based on newest, up to three of them with that label, so we need to page if there are more and still not getting to label.
// // Submits a GitHub issue that includes the debug logs and parsed stack trace.
// func create_github_issue(message vopr_message, output *vopr_output, issue_file_name string) error {
// 	body := create_issue_markdown(message, output)
// 	if output.seed_passed {
// 		body = "Note this seed passed when it was rerun by the VOPR Hub.<br><br>" + body
// 	}
// 	// Removes the file path from the name.
// 	issue_file_name = strings.Replace(issue_file_name, issue_directory+"/", "", 1)
// 	issue_contents := fmt.Sprintf(
// 		"{ \"title\": \"%s\", \"body\": \"%s\", \"labels\":[] }",
// 		issue_file_name,
// 		body,
// 	)
// 	issue := strings.NewReader(issue_contents)
// 	post_request, error := http.NewRequest(
// 		"POST",
// 		repository_url,
// 		issue,
// 	)
// 	if error != nil {
// 		log_error("Failed to create the HTTP request for the GitHub API")
// 		panic(error.Error())
// 	}
// 	// TODO: remove token
// 	post_request.Header.Set("Authorization", "token "+ "***")

// 	post_response, error := http.DefaultClient.Do(post_request)
// 	if error != nil {
// 		log_error("Failed to send the HTTP request for the GitHub API")
// 		panic(error.Error())
// 	}

// 	defer post_response.Body.Close()

// 	if post_response.StatusCode < 200 || post_response.StatusCode > 299 {
// 		error_message := fmt.Sprintf(
// 			"Received a non 2xx status code from GitHub. StatusCode: %d. Status: %s",
// 			post_response.StatusCode,
// 			post_response.Status,
// 		)
// 		log_error(error_message)
// 		return fmt.Errorf(error_message)
// 	}

// 	log_info(
// 		fmt.Sprintf(
// 			"GitHub issue has been created. Received StatusCode: %d and Status: %s.",
// 			post_response.StatusCode,
// 			post_response.Status,
// 		),
// 		message.hash[:],
// 	)
// 	return nil
// }

// // Creates the string that forms the body of the GitHub issue.
// func create_issue_markdown(message vopr_message, output *vopr_output) string {
// 	// Limit set here to avoid writing only a few characters for any particular section.
// 	const min_useful_length = 100

// 	// Extract the information about the conditions under which the VOPR runs TigerBeetle.
// 	output.extract_parameters(&message)

// 	length_of_stack_trace := len(output.stack_trace)
// 	length_of_parameters := len(output.parameters)
// 	length_of_logs := len(output.logs)
// 	remaining_space := max_github_issue_size

// 	stack_trace := ""
// 	parameters := ""
// 	debug_logs := ""

// 	// Set the stack trace string
// 	if length_of_stack_trace > 0 {
// 		if length_of_stack_trace <= max_github_issue_size {
// 			stack_trace = make_markdown_compatible(string(output.stack_trace[:]))
// 			remaining_space -= length_of_stack_trace
// 		} else {
// 			// If the stack trace is too large then just capture the beginning of it.
// 			stack_trace = make_markdown_compatible(
// 				string(output.stack_trace[:max_github_issue_size-4]),
// 			)
// 			stack_trace += "..."
// 			remaining_space = 0
// 		}
// 	}

// 	// Set the parameters string
// 	if length_of_parameters > 0 && remaining_space >= min_useful_length {
// 		if length_of_parameters <= remaining_space {
// 			parameters = make_markdown_compatible(output.parameters[:])
// 			remaining_space -= length_of_parameters
// 		} else {
// 			// If the parameters section is too large then just capture the beginning of it.
// 			parameters = make_markdown_compatible(output.parameters[:remaining_space-4]) + "..."
// 			remaining_space = 0
// 		}
// 	}

// 	// Set the debug logs string
// 	if remaining_space >= min_useful_length {
// 		if length_of_logs < remaining_space {
// 			debug_logs = make_markdown_compatible(string(output.logs[:]))
// 		} else {
// 			// Get the tail of the logs
// 			tail_start_index := length_of_logs - remaining_space
// 			debug_logs = make_markdown_compatible(string(output.logs[tail_start_index:]))
// 		}
// 	}

// 	var bugType string
// 	switch message.bug {
// 	case 1:
// 		bugType = "correctness"
// 	case 2:
// 		bugType = "liveness"
// 	case 3:
// 		bugType = "crash"
// 	default:
// 		panic("unreachable")
// 	}

// 	var body string
// 	if len(parameters) > 0 {
// 		body += fmt.Sprintf(
// 			"**Parameters:**<br>%s<br><br>",
// 			parameters,
// 		)
// 	}
// 	if len(stack_trace) > 0 {
// 		body += fmt.Sprintf(
// 			"**Stack Trace:**<br>%s<br><br>",
// 			stack_trace,
// 		)
// 	}
// 	if len(debug_logs) > 0 {
// 		body += fmt.Sprintf(
// 			"**Tail of Debug Logs:**<br>%s<br><br>",
// 			debug_logs,
// 		)
// 	}

// 	timestamp := time.Now().UTC().String()

// 	markdown := fmt.Sprintf(
// 		"**Bug Type:**<br>%s<br><br>**Seed:**<br>%d<br><br>**Commit Hash:**<br>%s<br><br>%s**Time:**<br>%s",
// 		bugType,
// 		message.seed,
// 		hex.EncodeToString(message.commit[:]),
// 		body,
// 		timestamp,
// 	)
// 	return markdown
// }

// // Escape characters that have special use in Markdown.
// func make_markdown_compatible(text string) string {
// 	text = strings.ReplaceAll(text, "\"", "\\\"")
// 	text = strings.ReplaceAll(text, "<", "&lt;")
// 	text = strings.ReplaceAll(text, ">", "&gt;")
// 	text = strings.ReplaceAll(text, "\n", "<br>")
// 	return text
// }
