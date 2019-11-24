# -----------------------------------------------------------------------------
# Stage: BUILD
# -----------------------------------------------------------------------------
FROM golang:1.13.4-alpine3.10 as BUILD

ARG BRANCH="master"

# Add github
RUN apk update && \
    apk upgrade && \
    apk add --no-cache \
        git

WORKDIR /go/src

RUN git clone https://github.com/bryan-nice/jira-issue-creation.git && \
    cd jira-issue-creation && \
    git checkout ${BRANCH}

RUN cd jira-issue-creation && \
    go mod download && \
    go get ./... && \
    go install ./...

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

ENTRYPOINT ["jira-issue-creation"]
