#~/bin/bash
tell application "iTerm"
    # etc...
    exec command "$@"

# chmod +x run-in-iterm.sh
# ./run-in-iterm.sh "echo 'hello world'