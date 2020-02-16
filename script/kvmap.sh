#!/bin/bash

declare -A SEARCH_REPLACE_MAP
function setup-kv-map
{
    f1="vgo-tools"
    r1="tools"
    # f2="search-2"
    # r2="replace-2"
    SEARCH_REPLACE_MAP=(
        [${f1}]=${r1}
        # [${f2}]=${r2}
    )
}

function list-kv-map
{
    if (( ! ${#} )); then
        Error1 function list-kv-map: No argument given
    fi
    local name="${1}"
    local -n KV_MAP="${name}"
    
    echo "+-------------------------------------------------------------------+"
    echo "|                      REPO REPAIR CHANGE TABLE                     |"
    echo "+-------------------------------------------------------------------+"
    printf "| %-32.32s| %-32.32s|\n" Search Replace
    echo "+-------------------------------------------------------------------+"

    for key in "${!KV_MAP[@]}"; do
        printf "| %-32.32s| %-32.32s|\n" "${key}" "${KV_MAP[${key}]}"
    done
    echo "+-------------------------------------------------------------------+"
}

