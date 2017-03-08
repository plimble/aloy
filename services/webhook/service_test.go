package webhook

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type WebhookServiceSuite struct {
	suite.Suite
	service *Service
}

func TestWebhookServiceSuite(t *testing.T) {
	suite.Run(t, &WebhookServiceSuite{})
}

func (t *WebhookServiceSuite) SetupTest() {
	t.service = New()
}

func (t *WebhookServiceSuite) TestParseGithubWebhook() {
	payload := []byte(`
	{
		"ref": "refs/heads/changes",
		"before": "9049f1265b7d61be4a8904a9a27120d2064dab3b",
		"after": "0d1a26e67d8f5eaf1f6ba5c57fc3c7d91ac0fd1c",
		"created": false,
		"deleted": false,
		"forced": false,
		"base_ref": null,
		"compare": "https://github.com/baxterthehacker/public-repo/compare/9049f1265b7d...0d1a26e67d8f",
		"commits": [
			{
				"id": "0d1a26e67d8f5eaf1f6ba5c57fc3c7d91ac0fd1c",
				"tree_id": "f9d2a07e9488b91af2641b26b9407fe22a451433",
				"distinct": true,
				"message": "Update README.md",
				"timestamp": "2015-05-05T19:40:15-04:00",
				"url": "https://github.com/baxterthehacker/public-repo/commit/0d1a26e67d8f5eaf1f6ba5c57fc3c7d91ac0fd1c",
				"author": {
					"name": "baxterthehacker",
					"email": "baxterthehacker@users.noreply.github.com",
					"username": "baxterthehacker"
				},
				"committer": {
					"name": "baxterthehacker",
					"email": "baxterthehacker@users.noreply.github.com",
					"username": "baxterthehacker"
				},
				"added": [

				],
				"removed": [

				],
				"modified": [
					"README.md"
				]
			}
		],
		"head_commit": {
			"id": "0d1a26e67d8f5eaf1f6ba5c57fc3c7d91ac0fd1c",
			"tree_id": "f9d2a07e9488b91af2641b26b9407fe22a451433",
			"distinct": true,
			"message": "Update README.md",
			"timestamp": "2015-05-05T19:40:15-04:00",
			"url": "https://github.com/baxterthehacker/public-repo/commit/0d1a26e67d8f5eaf1f6ba5c57fc3c7d91ac0fd1c",
			"author": {
				"name": "baxterthehacker",
				"email": "baxterthehacker@users.noreply.github.com",
				"username": "baxterthehacker"
			},
			"committer": {
				"name": "baxterthehacker",
				"email": "baxterthehacker@users.noreply.github.com",
				"username": "baxterthehacker"
			},
			"added": [

			],
			"removed": [

			],
			"modified": [
				"README.md"
			]
		},
		"repository": {
			"id": 35129377,
			"name": "public-repo",
			"full_name": "baxterthehacker/public-repo",
			"owner": {
				"name": "baxterthehacker",
				"email": "baxterthehacker@users.noreply.github.com"
			},
			"private": false,
			"html_url": "https://github.com/baxterthehacker/public-repo",
			"description": "desc",
			"fork": false,
			"url": "https://github.com/baxterthehacker/public-repo",
			"forks_url": "https://api.github.com/repos/baxterthehacker/public-repo/forks",
			"keys_url": "https://api.github.com/repos/baxterthehacker/public-repo/keys{/key_id}",
			"collaborators_url": "https://api.github.com/repos/baxterthehacker/public-repo/collaborators{/collaborator}",
			"teams_url": "https://api.github.com/repos/baxterthehacker/public-repo/teams",
			"hooks_url": "https://api.github.com/repos/baxterthehacker/public-repo/hooks",
			"issue_events_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/events{/number}",
			"events_url": "https://api.github.com/repos/baxterthehacker/public-repo/events",
			"assignees_url": "https://api.github.com/repos/baxterthehacker/public-repo/assignees{/user}",
			"branches_url": "https://api.github.com/repos/baxterthehacker/public-repo/branches{/branch}",
			"tags_url": "https://api.github.com/repos/baxterthehacker/public-repo/tags",
			"blobs_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/blobs{/sha}",
			"git_tags_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/tags{/sha}",
			"git_refs_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/refs{/sha}",
			"trees_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/trees{/sha}",
			"statuses_url": "https://api.github.com/repos/baxterthehacker/public-repo/statuses/{sha}",
			"languages_url": "https://api.github.com/repos/baxterthehacker/public-repo/languages",
			"stargazers_url": "https://api.github.com/repos/baxterthehacker/public-repo/stargazers",
			"contributors_url": "https://api.github.com/repos/baxterthehacker/public-repo/contributors",
			"subscribers_url": "https://api.github.com/repos/baxterthehacker/public-repo/subscribers",
			"subscription_url": "https://api.github.com/repos/baxterthehacker/public-repo/subscription",
			"commits_url": "https://api.github.com/repos/baxterthehacker/public-repo/commits{/sha}",
			"git_commits_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/commits{/sha}",
			"comments_url": "https://api.github.com/repos/baxterthehacker/public-repo/comments{/number}",
			"issue_comment_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/comments{/number}",
			"contents_url": "https://api.github.com/repos/baxterthehacker/public-repo/contents/{+path}",
			"compare_url": "https://api.github.com/repos/baxterthehacker/public-repo/compare/{base}...{head}",
			"merges_url": "https://api.github.com/repos/baxterthehacker/public-repo/merges",
			"archive_url": "https://api.github.com/repos/baxterthehacker/public-repo/{archive_format}{/ref}",
			"downloads_url": "https://api.github.com/repos/baxterthehacker/public-repo/downloads",
			"issues_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues{/number}",
			"pulls_url": "https://api.github.com/repos/baxterthehacker/public-repo/pulls{/number}",
			"milestones_url": "https://api.github.com/repos/baxterthehacker/public-repo/milestones{/number}",
			"notifications_url": "https://api.github.com/repos/baxterthehacker/public-repo/notifications{?since,all,participating}",
			"labels_url": "https://api.github.com/repos/baxterthehacker/public-repo/labels{/name}",
			"releases_url": "https://api.github.com/repos/baxterthehacker/public-repo/releases{/id}",
			"created_at": 1430869212,
			"updated_at": "2015-05-05T23:40:12Z",
			"pushed_at": 1430869217,
			"git_url": "git://github.com/baxterthehacker/public-repo.git",
			"ssh_url": "git@github.com:baxterthehacker/public-repo.git",
			"clone_url": "https://github.com/baxterthehacker/public-repo.git",
			"svn_url": "https://github.com/baxterthehacker/public-repo",
			"homepage": null,
			"size": 0,
			"stargazers_count": 0,
			"watchers_count": 0,
			"language": null,
			"has_issues": true,
			"has_downloads": true,
			"has_wiki": true,
			"has_pages": true,
			"forks_count": 0,
			"mirror_url": null,
			"open_issues_count": 0,
			"forks": 0,
			"open_issues": 0,
			"watchers": 0,
			"default_branch": "master",
			"stargazers": 0,
			"master_branch": "master"
		},
		"pusher": {
			"name": "baxterthehacker",
			"email": "baxterthehacker@users.noreply.github.com"
		},
		"sender": {
			"login": "baxterthehacker",
			"id": 6752317,
			"avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
			"gravatar_id": "",
			"url": "https://api.github.com/users/baxterthehacker",
			"html_url": "https://github.com/baxterthehacker",
			"followers_url": "https://api.github.com/users/baxterthehacker/followers",
			"following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
			"gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
			"starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
			"subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
			"organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
			"repos_url": "https://api.github.com/users/baxterthehacker/repos",
			"events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
			"received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
			"type": "User",
			"site_admin": false
		}
	}
	`)
	expWebhook := &Webhook{
		SenderName:     "baxterthehacker",
		SenderAvatar:   "https://avatars.githubusercontent.com/u/6752317?v=3",
		Commit:         "0d1a26e67d8f5eaf1f6ba5c57fc3c7d91ac0fd1c",
		Ref:            "refs/heads/changes",
		RepoName:       "baxterthehacker/public-repo",
		RepoHomepage:   "https://github.com/baxterthehacker/public-repo",
		RepoDesciption: "desc",
		HTTPURL:        "https://github.com/baxterthehacker/public-repo.git",
	}

	webhook, err := t.service.ParseGithubWebhook(payload)
	t.NoError(err)
	t.Equal(expWebhook, webhook)
}

func (t *WebhookServiceSuite) TestParseGitlabWebhook() {
	payload := []byte(`
	{
		"object_kind": "push",
		"before": "95790bf891e76fee5e1747ab589903a6a1f80f22",
		"after": "da1560886d4f094c3e6c9ef40349f7d38b5d27d7",
		"ref": "refs/heads/master",
		"checkout_sha": "da1560886d4f094c3e6c9ef40349f7d38b5d27d7",
		"user_id": 4,
		"user_name": "John Smith",
		"user_email": "john@example.com",
		"user_avatar": "https://s.gravatar.com/avatar/d4c74594d841139328695756648b6bd6?s=8://s.gravatar.com/avatar/d4c74594d841139328695756648b6bd6?s=80",
		"project_id": 15,
		"project":{
			"name":"Diaspora",
			"description":"desc",
			"web_url":"http://example.com/mike/diaspora",
			"avatar_url":null,
			"git_ssh_url":"git@example.com:mike/diaspora.git",
			"git_http_url":"http://example.com/mike/diaspora.git",
			"namespace":"Mike",
			"visibility_level":0,
			"path_with_namespace":"mike/diaspora",
			"default_branch":"master",
			"homepage":"http://example.com/mike/diaspora",
			"url":"git@example.com:mike/diaspora.git",
			"ssh_url":"git@example.com:mike/diaspora.git",
			"http_url":"http://example.com/mike/diaspora.git"
		},
		"repository":{
			"name": "Diaspora",
			"url": "git@example.com:mike/diaspora.git",
			"description": "",
			"homepage": "http://example.com/mike/diaspora",
			"git_http_url":"http://example.com/mike/diaspora.git",
			"git_ssh_url":"git@example.com:mike/diaspora.git",
			"visibility_level":0
		},
		"commits": [
			{
				"id": "b6568db1bc1dcd7f8b4d5a946b0b91f9dacd7327",
				"message": "Update Catalan translation to e38cb41.",
				"timestamp": "2011-12-12T14:27:31+02:00",
				"url": "http://example.com/mike/diaspora/commit/b6568db1bc1dcd7f8b4d5a946b0b91f9dacd7327",
				"author": {
					"name": "Jordi Mallach",
					"email": "jordi@softcatala.org"
				},
				"added": ["CHANGELOG"],
				"modified": ["app/controller/application.rb"],
				"removed": []
			},
			{
				"id": "da1560886d4f094c3e6c9ef40349f7d38b5d27d7",
				"message": "fixed readme",
				"timestamp": "2012-01-03T23:36:29+02:00",
				"url": "http://example.com/mike/diaspora/commit/da1560886d4f094c3e6c9ef40349f7d38b5d27d7",
				"author": {
					"name": "GitLab dev user",
					"email": "gitlabdev@dv6700.(none)"
				},
				"added": ["CHANGELOG"],
				"modified": ["app/controller/application.rb"],
				"removed": []
			}
		],
		"total_commits_count": 4
	}
	`)
	expWebhook := &Webhook{
		SenderName:     "John Smith",
		SenderAvatar:   "https://s.gravatar.com/avatar/d4c74594d841139328695756648b6bd6?s=8://s.gravatar.com/avatar/d4c74594d841139328695756648b6bd6?s=80",
		Commit:         "da1560886d4f094c3e6c9ef40349f7d38b5d27d7",
		Ref:            "refs/heads/master",
		RepoName:       "mike/diaspora",
		RepoHomepage:   "http://example.com/mike/diaspora",
		RepoDesciption: "desc",
		HTTPURL:        "http://example.com/mike/diaspora.git",
	}

	webhook, err := t.service.ParseGitlabWebhook(payload)
	t.NoError(err)
	t.Equal(expWebhook, webhook)
}
