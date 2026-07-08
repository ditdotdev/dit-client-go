schema_version = 1

project {
  license          = "BUSL-1.1"
  copyright_holder = "Dit"
  copyright_year   = 2026

  header_ignore = [
    "gradlew",
    "gradlew.bat",
    "gradle/**",
    "build/**",
    "**/build/**",
    ".health/**",
    "**/*.out",
    # OpenAPI-generated (regenerated from the server spec each release):
    "api_commits.go",
    "api_contexts.go",
    "api_operations.go",
    "api_remotes.go",
    "api_repositories.go",
    "api_volumes.go",
    "client.go",
    "configuration.go",
    "model_api_error.go",
    "model_commit.go",
    "model_commit_status.go",
    "model_context.go",
    "model_operation.go",
    "model_progress_entry.go",
    "model_remote.go",
    "model_remote_parameters.go",
    "model_repository.go",
    "model_repository_status.go",
    "model_volume.go",
    "model_volume_status.go",
    "response.go",
    "utils.go",
  ]
}
