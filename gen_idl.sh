#!/bin/bash
set -e

# run demo command like
# ./gen_idl.sh pb_gen idl
# ./gen_idl.sh pb_gen/user/ idl/user/


# $1 生成文件路径
# $2 .proto 文件路径
function scan() {
  list=$(ls "${2}")
  for f in $list; do
    if test -d "$2"/"$f"; then
      scan "$1"/"$f" "$2"/"$f"
    else
      mkdir -p "$1"
      protoc --proto_path="$2" -I="idl" --go_opt=Mpublic/public.proto=github.com/twelveeee/amis-admin-go/pb_gen/public  --go_out="$1"  "$2"/"$f"
      go_file=$1/${f/proto/pb.go}
      if [ `uname -s` = "Darwin" ]; then
        /usr/bin/sed -i "" -e 's/,omitempty//g' "$go_file";
      else
        sed -i 's/,omitempty//g' "$go_file";
      fi
      echo "$go_file"
    fi
  done
}

if [ $# == 0 ]; then
  scan pb_gen idl
else
  scan "$1" "$2"
fi

