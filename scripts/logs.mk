#####################
##@ Logs            
#####################

log: ## log env vars
	@echo "\n======================================================\n"
	@echo "PROJECT_ROOT         $(PROJECT_ROOT)"
	@echo "REPO_URL             $(REPO_URL)"
	@echo "COMMIT               $(COMMIT)"
	@echo "BRANCH               $(BRANCH)"
	@echo "APP_NAME             $(APP_NAME)"
	@echo "DATE                 $(DATE)"
	@echo "\n======================================================\n"


# HELP
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help: ## This help.
	@awk 'BEGIN {FS = ":.*##"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
