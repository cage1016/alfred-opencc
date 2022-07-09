#!/usr/bin/env bash

# Workflow environment variables
# These variables create an Alfred-like environment

root="$( git rev-parse --show-toplevel )"
testdir="${root}/testenv"

# Absolute bare-minimum for AwGo to function...
export alfred_workflow_bundleid="com.kaichu.opencc"
export alfred_workflow_data="${testdir}/data"
export alfred_workflow_cache="${testdir}/cache"

test -f "$HOME/Library/Preferences/com.runningwithcrayons.Alfred.plist" || {
	export alfred_version="3.8.1"
}

# Expected by ExampleNew
export alfred_workflow_version="0.1.0"
export alfred_workflow_name="Open Chinese Convert"
export alfred_workflow_package="github.com/cage1016/alfred-opencc"
# export alfred_workflow_category="Tools" // Internet/Tools/Productivity or empty
export alfred_workflow_description="Convert Chinese between Traditional and Simplified by OpenCC"
export alfred_workflow_created_by="KAI CHU CHUNG"
export alfred_workflow_website="https://kaichu.io"

# Prevent random ID from being generated
export AW_SESSION_ID="test-session-id"

export GO111MODULE=on