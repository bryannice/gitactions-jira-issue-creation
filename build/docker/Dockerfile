# -----------------------------------------------------------------------------
# Stage: BUILD
# -----------------------------------------------------------------------------
FROM golang:1.13.4-alpine3.10 as BUILD

# Add github
RUN apk update && \
    apk upgrade && \
    apk add --no-cache \
        git \
        make

WORKDIR /go/src

# Clone repo
RUN git clone https://github.com/bryan-nice/jira-issue-creation.git

# Checkout to target branch and if not master, then locally test
# merging master into branch.
WORKDIR /go/src/jira-issue-creation

# Building go app
RUN make build

# -----------------------------------------------------------------------------
# Stage: BUILD
# -----------------------------------------------------------------------------
FROM alpine:3.10

LABEL "com.github.actions.icon"="message-square"
LABEL "com.github.actions.color"="purple"
LABEL "com.github.actions.name"="Jira Issue Creation"
LABEL "com.github.actions.description"="Create a Jira Issue"

COPY --from=BUILD /go/bin/jira-issue-creation /usr/bin
COPY --from=BUILD /go/src/jira-issue-creation/LICENSE /
COPY --from=BUILD /go/src/jira-issue-creation/README.md /

RUN mkdir -p /github

WORKDIR /github

ENTRYPOINT ["jira-issue-creation"]