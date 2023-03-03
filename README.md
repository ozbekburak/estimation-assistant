# AI Based Task Estimation

## Introduction
It's crucial to deliver projects on time and within budget. However, estimating the time and resources needed to complete a task can be a challenging task, especially when it involves multiple factors such as complexity, skill level, resources, and risks. To address this challenge, I have created a Go-based task estimation application that can help product managers and engineers to estimate the time and resources required to complete a task accurately.


## Features
The task estimation application will have the following features:

- The application will prompt the user to answer several questions related to the task they need to work on. These questions will cover the following aspects of the task:
- Task Description
- Task Complexity
- Skill Level
- Resources
- Timeline
- Risks and Issues

**Prompt Creation:** Based on the user's input, the application will create a complete prompt that can be used in an AI-based application to estimate the time and resources needed to complete the task accurately.

**Time Estimation:** The prompt created by the application can be used in conjunction with an AI-based application to estimate the time required to complete the task accurately. The AI-based application can consider various factors, including complexity, skill level, and resources, to provide an accurate estimate.

**Pre-prepared Prompt:** The application can also be used with the prepared prompt to estimate the time and resources required to complete the task accurately. The feature of using a pre-prepared prompt with the **"-prompt"** argument can be beneficial in situations where there is an error in the application (because it is saving to file before asking to ChatGPT), or the user needs to use a prompt from their storage. 

- Suppose the user has previously created a prompt for a similar task and stored it on their system. In that case, they can use the "-prompt" argument to load the prompt directly into the program, saving time and effort. 


## Usage

User needs to set the [OPENAI API KEY](https://platform.openai.com/account/api-keys) environment variable. After that:

```go
export OPENAI_API_KEY={your_api_key}
git clone github.com/ozbekburak/estimation-assistant
cd estimation-assistant
go run .
```
or

```go
export OPENAI_API_KEY={your_api_key}
git clone github.com/ozbekburak/estimation-assistant
cd estimation-assistant
go build .
./estimation-assistant -prompt "./prompts/prompt-20230303165515.txt"
```

## Screenshots

### Prompt Creating with User Inputs
![ask-to-user](https://github.com/ozbekburak/dfir-radar/blob/main/img/ask-to-user.png?raw=true)

### Using Pre-Prepared Prompt
![preprepared](https://github.com/ozbekburak/dfir-radar/blob/main/img/preprepared.png?raw=true)

## Conclusion

The task estimation application can provide a valuable tool for product managers and engineers to create better plans and make informed decisions. However, it's essential to remember that the estimates provided by the application are just that, estimates, and not the final word on the time and resources required to complete a task.