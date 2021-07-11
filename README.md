# better-branch
Interactive CLI helper for creating git branches with JIRA Links and some text

## Is this a buggy beta?

Yes

## How it works?
![Enter JIRA link and/or text, then create a new branch from it](https://media.giphy.com/media/ji7D1GBEFgQRQE0oHM/giphy.gif)

There are multiple ways to create branch name with Jira and/or text description.
Currenly new branch is checkouted after creation by default.

### Interactive
Just launch without any args. When asked for values, enter them. If name is beautiful for you, create a new branch.

### Pass arguments
```shell
  -c true
    	Checkout to new branch (default true (default true)
  -i string
    	JIRA Link or issue
  -t string
    	Custom Issue Text
```
Arguments could be passed with keywords or shorthand.

```shell
./main -i https://some.jira.cloud/issues/ABC-123 -t Add big button

# or
./main https://some.jira.cloud/issues/ABC-123  Add big button

# or even
./main https://some.jira.cloud/issues/ABC-123
```


## OS support
Currently only OSX and Unix are supported

## TODO:
- [ ] Go back, add more text, and other interactive mode impovements
- [ ] Set autocheckout true/false with interactive mode
- [ ] Support prefixes like `feature/`
- [ ] Support other popular issue trackers like YouTrack, Asana, etc
- [ ] Save screen space in interactive mode by putting hints onto background
