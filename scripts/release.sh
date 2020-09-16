#!/bin/sh
bump_type=$1
current_version=$(cat ../version)
flavours=$(ls -d ../*/ | awk '!/scripts/')
next_version=$(./semver.sh $bump_type $current_version)
origin=$(git remote get-url origin)

echo "$current_version ==> $next_version"

release() {
    dir=$1
    flavour=${dir:3}
    flavour=${flavour%?}
    tagname=v$next_version-$flavour
    echo releasing $tagname...
    cd $dir

    git init
    git remote add origin $origin
    git add . > /dev/null
    git commit -m "release/$next_version" > /dev/null
    git tag $tagname > /dev/null
    git push origin $tagname
    rm -rf .git
}

for flavour in $flavours
do
    release $flavour
done

echo $next_version > ../version