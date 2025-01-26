install_precommit:
	@pip install pre-commit==3.5.0
	@pre-commit install --install-hooks  # Install the defined hooks in the .pre-commit-config.yaml
	@pre-commit install --hook-type commit-msg  # Install the commit-msg hook for verifying commit messages

update_precommit:
	@echo "Updating pre-commit hooks..."
	@pre-commit autoupdate  # Update all hooks to the latest compatible versions

run_precommit:
	@echo "Running pre-commit hooks..."
	@pre-commit run --all-files  # Run all the pre-commit hooks on all files

run_build_project:
	@echo "Building the project..."

start_project:
	@echo "Starting the project..."

