FROM migrate/migrate

COPY ./migrations                               ./migrations
COPY ./docker/migrations/scripts/migration.sh   ./migration.sh


ENTRYPOINT ["./migration.sh"]
