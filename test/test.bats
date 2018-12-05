#!/usr/bin/env bats


if [ ! -x "$(dirname $BATS_TEST_DIRNAME)/t3" ]; then
    printf "No executable ../t3 found.\n" >&2
    printf "Run \`script/build\` in the top-level and try again.\n" >&2
    exit 1
fi

# include t3 in PATH
PATH=/usr/bin:/bin:/usr/sbin:/sbin
PATH="$(dirname $BATS_TEST_DIRNAME):$PATH"

@test "sanity" {
    run true
    [ "$status" -eq 0 ]
    [ "$output" = "" ]
}


@test "spacecat empty file" {
    run t3 -d $BATS_TEST_DIRNAME/data.json -t $BATS_TEST_DIRNAME/template_empty.mustache
    [ "$status" -eq 0 ]
    [ "$output" = "" ]
}

@test "spacecat basic file" {
    run t3 -d $BATS_TEST_DIRNAME/data.json -t $BATS_TEST_DIRNAME/template_basic.mustache
    [ "$output" = "hello world" ]
    echo "$output"
    [ "$status" -eq 0 ]
}
