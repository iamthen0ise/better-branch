# bb (better-branch)
Interactive CLI helper for creating git branches with JIRA Links and some text

## Still in development?
Yes

## How it works?
[![asciicast](https://asciinema.org/a/d4NPyH679pdgjJVfLQpV4SOf7.svg)](https://asciinema.org/a/d4NPyH679pdgjJVfLQpV4SOf7)
This tiny utility was made when i completely bored of creating JIRA branches on Web interface and pulling it to local.

Just call `bb`, then Enter JIRA link and/or text, then create a new branch from.

There are multiple ways to create branch name with Jira and/or text description.
### Interactive
Just launch without any args. When asked for values, enter them. If name is beautiful for you, create a new branch.

### Pass arguments
```shell
  -f feature
        Create `feature/*` branch
  -h hotfix
        Create `hotfix/*` branch
  -c true
    	Checkout to new branch (default true (default true)
  -i string
    	JIRA Link or issue
  -t string
    	Custom Issue Text
```
Arguments could be passed with keywords or shorthand.

```shell
./main -f -i https://some.jira.cloud/issues/ABC-123 -t Add big button

# or
./main f https://some.jira.cloud/issues/ABC-123  Add big button

# or even
./main https://some.jira.cloud/issues/ABC-123
```

New branch is checkouted after creation by default.

## OS support
Builds are made for Windows, OSX and Linux by Goreleaser. But code wasn't tested on Linux and Windows.

## TODO:
- [ ] Go back, add more text, and other interactive mode impovements
- [ ] Set autocheckout true/false with interactive mode
- [x] Support prefixes like `feature/`
- [ ] Support other popular issue trackers like YouTrack, Asana, etc
- [ ] Save screen space in interactive mode by putting hints onto background
