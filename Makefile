.PHONY: setup_arts_folder clean_arts_folder

ARTS_FOLDER:="ARTS"
ARTS_DATE_DIR:=$(ARTS_FOLDER)/`date "+%Y%m%d"`
ARTS_ALGORITHM_FILE:=$(ARTS_DATE_DIR)/algorithm.md
ARTS_REVIEW_FILE:=$(ARTS_DATE_DIR)/review.md
ARTS_TIP_FILE:=$(ARTS_DATE_DIR)/tip.md
ARTS_SHARE_FILE:=$(ARTS_DATE_DIR)/share.md

setup_arts_folder: clean_arts_folder
	@cd $(ARTS_FOLDER)
	@mkdir $(ARTS_DATE_DIR)
	@touch $(ARTS_ALGORITHM_FILE)
	@touch $(ARTS_REVIEW_FILE)
	@touch $(ARTS_TIP_FILE)
	@touch $(ARTS_SHARE_FILE)
	@echo "[done] Setup ARTS folder for $(ARTS_DATE_DIR) "
	@git add .
	@git status

clean_arts_folder:
	@rm -rf $(ARTS_DATE_DIR)

setup_db:
	docker-compose up -d --remove-orphans
	docker-compose ps
	@echo "psql --host localhost --port 55432 -U postgres"
	@echo "mysql -h 127.0.0.1 -P 33060 -u root --ssl-mode DISABLED -p" # empty password

clean:
	find . -name a.out | xargs -I {} rm {}
