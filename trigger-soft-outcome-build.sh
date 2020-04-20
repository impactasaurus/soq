# This is based on https://github.com/mernst/plume-lib/blob/master/bin/trigger-travis.sh
# It triggers a rebuild of the soft outcomes project
# the only requirement is a TRAVIS_ACCESS_TOKEN env var, the token needs permissions to trigger soft outcome builds

TRAVIS_URL=travis-ci.org
BRANCH=master
USER=impactasaurus
REPO=softoutcomes
TOKEN=$TRAVIS_ACCESS_TOKEN

if [ -n "$TRAVIS_REPO_SLUG" ] ; then
    MESSAGE=",\"message\": \"Triggered by upstream build of $TRAVIS_REPO_SLUG commit "`git rev-parse --short HEAD`"\""
else
    MESSAGE=""
fi

## For debugging:
echo "USER=$USER"
echo "REPO=$REPO"
echo "TOKEN=$TOKEN"
echo "MESSAGE=$MESSAGE"

body="{
\"request\": {
  \"branch\":\"$BRANCH\"
  $MESSAGE
}}"

curl -s -X POST \
  -H "Content-Type: application/json" \
  -H "Accept: application/json" \
  -H "Travis-API-Version: 3" \
  -H "Authorization: token ${TOKEN}" \
  -d "$body" \
  https://api.${TRAVIS_URL}/repo/${USER}%2F${REPO}/requests \
 | tee /tmp/travis-request-output.$$.txt

if grep -q '"@type": "error"' /tmp/travis-request-output.$$.txt; then
    exit 1
fi
if grep -q 'access denied' /tmp/travis-request-output.$$.txt; then
    exit 1
fi
