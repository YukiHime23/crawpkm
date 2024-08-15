# lunar vim
set -Ua fish_user_paths /home/kiennt1/.local/bin

# neo vim
set -Ua fish_user_paths /opt/nvim-linux64/bin 

# golang
set -Ua fish_user_paths /usr/local/go/bin
## protoc-gen-go
set -Ua fish_user_paths /home/kiennt1/go/bin

# bun & npm
set PATH /home/kiennt1/.nvm/versions/node/v20.12.2/bin $PATH
#set PATH /home/kiennt1/.local/share/nvm/v20.12.2/bin $PATH

# dotnet
set DOTNET_ROOT $HOME/.dotnet $DOTNET_ROOT
set PATH $PATH:$DOTNET_ROOT:$DOTNET_ROOT/tools $PATH

# set PATH <some-path> $PATH
