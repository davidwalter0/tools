#!/bin/bash
# https://misc.flogisoft.com/bash/tip_colors_and_formatting
# https://stackoverflow.com/a/28938235/1470495

READY_FILE=.${0##*/}.ready
function validate_ready
{
if ((! DRY_RUN )); then
    Warning2 DRY_RUN is off
    if [[ ! -e ${READY_FILE} ]]; then
        Error1 "${READY_FILE} not found.\n\n\ttouch ${READY_FILE} when testing is complete\n"
    fi
fi

}
function validate_arguments
{
    for arg in ${@}; do
        case ${arg} in
            --dry-run)
                Warning1 ${arg} sets DRY_RUN=1 true
                Info1 This will not modify git repo
                export DRY_RUN=1
                ;;
            --exec)
                Warning1 ${arg} sets DRY_RUN=0 false
                Warning1 May modify the git repo
                export DRY_RUN=0
                ;;
            *)
                Warning1 "Usage ${0##*/} [--dry-run|--exec]"
                Error1 No argument given
                ;;
        esac
    done
}

function validate_git_repo
{
    if root=$(git rev-parse --show-toplevel 2>/dev/null); then
        Info1 git repo root: ${root}
    else
        Error1 This path ${PWD} is not in a git repo.
    fi

    Info1 The top of this git repo is ${root}
}

function validate_git_find_grep
{
# This utility uses the git alias find-grep
# add the following to ~/.gitconfig to use
# [alias]
#         find-grep = "!git rev-list --all | xargs git grep "

    # verify find-grep exists

    if git find-grep 2>&1|grep 'is not a git command'; then
        cat <<EOF

# This utility uses the git alias find-grep
# add the following to ~/.gitconfig to use
[alias]
        find-grep = "!git rev-list --all | xargs git grep "
EOF
        Error1 ${0##*/} uses git alias 'find-grep' in .gitconfig
    fi

}
declare -A COLORMAP=(
    # Color_Off='\033[0m'
    [ON_BLACK]='\033[40m'
    [ON_RED]='\033[41m'
    [ON_GREEN]='\033[42m'
    [ON_YELLOW]='\033[43m'
    [ON_BLUE]='\033[44m'
    [ON_PURPLE]='\033[45m'
    [ON_CYAN]='\033[46m'
    [ON_WHITE]='\033[47m'
    [RESET]='\033[0m'
    [BBlack]='\033[1;30m'
    [BRed]='\033[1;31m'
    [BGreen]='\033[1;32m'
    [BYellow]='\033[1;33m'
    [BBlue]='\033[1;34m'
    [BPurple]='\033[1;35m'
    [BCyan]='\033[1;36m'
    [BWhite]='\033[1;37m'
)

function print-colormap
{
    for key in "${!COLORMAP[@]}"; do
        printf "[%-20s] " ${key}
        echo -e "${COLORMAP[${key}]}" "\${COLORMAP[${key}]}${COLORMAP[RESET]}"
    done
}
function log-print
{
    printf "$(date-stamp) [${0##*/}] ${*}\n"
}

function date_stamp
{
    date +%Y.%m.%d.%H.%M.%S
}

function Warning1
{
    echo -e "${COLORMAP[BRed]}$(date_stamp) Warning${COLORMAP[RESET]} ${*}"
}

function Warning2
{
    echo -e "${COLORMAP[ON_RED]}$(date_stamp) Warning${COLORMAP[RESET]} ${*}"
}

function Error1
{
    echo -e "${COLORMAP[ON_RED]}$(date_stamp) Error  ${COLORMAP[RESET]} ${*}\nExit"
    exit 1
}

function Info1
{
    echo -e "${COLORMAP[BBlue]}$(date_stamp) Info${COLORMAP[RESET]}    ${*}"
}

function Info2
{
    echo -e "${COLORMAP[ON_BLUE]}$(date_stamp) Info${COLORMAP[RESET]}    ${*}"
}


# echo -e '\033[2K'  # clear the screen and do not move the position
# or:

# echo -e '\033[2J\033[u' # clear the screen and reset the position

# # RESET
# Color_Off='\033[0m'       # Text RESET

# # Regular Colors
# Black='\033[0;30m'        # Black
# Red='\033[0;31m'          # Red
# Green='\033[0;32m'        # Green
# Yellow='\033[0;33m'       # Yellow
# Blue='\033[0;34m'         # Blue
# Purple='\033[0;35m'       # Purple
# Cyan='\033[0;36m'         # Cyan
# White='\033[0;37m'        # White

# # Bold
# BBlack='\033[1;30m'       # Black
# BRed='\033[1;31m'         # Red
# BGreen='\033[1;32m'       # Green
# BYellow='\033[1;33m'      # Yellow
# BBlue='\033[1;34m'        # Blue
# BPurple='\033[1;35m'      # Purple
# BCyan='\033[1;36m'        # Cyan
# BWhite='\033[1;37m'       # White

# # Underline
# UBlack='\033[4;30m'       # Black
# URed='\033[4;31m'         # Red
# UGreen='\033[4;32m'       # Green
# UYellow='\033[4;33m'      # Yellow
# UBlue='\033[4;34m'        # Blue
# UPurple='\033[4;35m'      # Purple
# UCyan='\033[4;36m'        # Cyan
# UWhite='\033[4;37m'       # White

# # Background
# On_Black='\033[40m'       # Black
# On_Red='\033[41m'         # Red
# On_Green='\033[42m'       # Green
# On_Yellow='\033[43m'      # Yellow
# On_Blue='\033[44m'        # Blue
# On_Purple='\033[45m'      # Purple
# On_Cyan='\033[46m'        # Cyan
# On_White='\033[47m'       # White

# # High Intensity
# IBlack='\033[0;90m'       # Black
# IRed='\033[0;91m'         # Red
# IGreen='\033[0;92m'       # Green
# IYellow='\033[0;93m'      # Yellow
# IBlue='\033[0;94m'        # Blue
# IPurple='\033[0;95m'      # Purple
# ICyan='\033[0;96m'        # Cyan
# IWhite='\033[0;97m'       # White

# # Bold High Intensity
# BIBlack='\033[1;90m'      # Black
# BIRed='\033[1;91m'        # Red
# BIGreen='\033[1;92m'      # Green
# BIYellow='\033[1;93m'     # Yellow
# BIBlue='\033[1;94m'       # Blue
# BIPurple='\033[1;95m'     # Purple
# BICyan='\033[1;96m'       # Cyan
# BIWhite='\033[1;97m'      # White

# # High Intensity backgrounds
# On_IBlack='\033[0;100m'   # Black
# On_IRed='\033[0;101m'     # Red
# On_IGreen='\033[0;102m'   # Green
# On_IYellow='\033[0;103m'  # Yellow
# On_IBlue='\033[0;104m'    # Blue
# On_IPurple='\033[0;105m'  # Purple
# On_ICyan='\033[0;106m'    # Cyan
# On_IWhite='\033[0;107m'   # White

# +------------+----------+---------+-------+------------------+------------------------------+--------------------------------------+
# | color-mode | octal    | hex     | bash  | description      | example (= in octal)         | NOTE                                 |
# +------------+----------+---------+-------+------------------+------------------------------+--------------------------------------+
# |          0 | \033[0m  | \x1b[0m | \e[0m | reset any affect | echo -e "\033[0m"            | 0m equals to m                       |
# |          1 | \033[1m  |         |       | light (= bright) | echo -e "\033[1m####\033[m"  | -                                    |
# |          2 | \033[2m  |         |       | dark (= fade)    | echo -e "\033[2m####\033[m"  | -                                    |
# +------------+----------+---------+-------+------------------+------------------------------+--------------------------------------+
# |  text-mode | ~        |         |       | ~                | ~                            | ~                                    |
# +------------+----------+---------+-------+------------------+------------------------------+--------------------------------------+
# |          3 | \033[3m  |         |       | italic           | echo -e "\033[3m####\033[m"  |                                      |
# |          4 | \033[4m  |         |       | underline        | echo -e "\033[4m####\033[m"  |                                      |
# |          5 | \033[5m  |         |       | blink (slow)     | echo -e "\033[3m####\033[m"  |                                      |
# |          6 | \033[6m  |         |       | blink (fast)     | ?                            | not wildly support                   |
# |          7 | \003[7m  |         |       | reverse          | echo -e "\033[7m####\033[m"  | it affects the background/foreground |
# |          8 | \033[8m  |         |       | hide             | echo -e "\033[8m####\033[m"  | it affects the background/foreground |
# |          9 | \033[9m  |         |       | cross            | echo -e "\033[9m####\033[m"  |                                      |
# +------------+----------+---------+-------+------------------+------------------------------+--------------------------------------+
# | foreground | ~        |         |       | ~                | ~                            | ~                                    |
# +------------+----------+---------+-------+------------------+------------------------------+--------------------------------------+
# |         30 | \033[30m |         |       | black            | echo -e "\033[30m####\033[m" |                                      |
# |         31 | \033[31m |         |       | red              | echo -e "\033[31m####\033[m" |                                      |
# |         32 | \033[32m |         |       | green            | echo -e "\033[32m####\033[m" |                                      |
# |         33 | \033[32m |         |       | yellow           | echo -e "\033[33m####\033[m" |                                      |
# |         34 | \033[32m |         |       | blue             | echo -e "\033[34m####\033[m" |                                      |
# |         35 | \033[32m |         |       | purple           | echo -e "\033[35m####\033[m" | real name: magenta = reddish-purple  |
# |         36 | \033[32m |         |       | cyan             | echo -e "\033[36m####\033[m" |                                      |
# |         37 | \033[32m |         |       | white            | echo -e "\033[37m####\033[m" |                                      |
# +------------+----------+---------+-------+------------------+------------------------------+--------------------------------------+
# |         38 | 8/24     |                    This is for special use of 8-bit or 24-bit                                            |
# +------------+----------+---------+-------+------------------+------------------------------+--------------------------------------+
# | background | ~        |         |       | ~                | ~                            | ~                                    |
# +------------+----------+---------+-------+------------------+------------------------------+--------------------------------------+
# |         40 | \033[40m |         |       | black            | echo -e "\033[40m####\033[m" |                                      |
# |         41 | \033[41m |         |       | red              | echo -e "\033[41m####\033[m" |                                      |
# |         42 | \033[42m |         |       | green            | echo -e "\033[42m####\033[m" |                                      |
# |         43 | \033[43m |         |       | yellow           | echo -e "\033[43m####\033[m" |                                      |
# |         44 | \033[44m |         |       | blue             | echo -e "\033[44m####\033[m" |                                      |
# |         45 | \033[45m |         |       | purple           | echo -e "\033[45m####\033[m" | real name: magenta = reddish-purple  |
# |         46 | \033[46m |         |       | cyan             | echo -e "\033[46m####\033[m" |                                      |
# |         47 | \033[47m |         |       | white            | echo -e "\033[47m####\033[m" |                                      |
# +------------+----------+---------+-------+------------------+------------------------------+--------------------------------------+
# |         48 | 8/24     |                    This is for special use of 8-bit or 24-bit                                            |
# +------------+----------+---------+-------+------------------+------------------------------+--------------------------------------+
