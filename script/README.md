#### Repository text replacement utility

Clean up vestigial, secret or out-of-date inconsistent information

**_This isn't for daily use as rewriting git history is generally bad
practice for public or shared repositories_**

Backup your repository prior to executing this command.

**_This will invalidate references to this repository's history_**

**_This command rewrites the git repository history_**


Replace old text (map key) with new text (map value) by updating te
kvmap.sh

Process to use and test

1. Backup your repo
2. Backup your repo
3. Backup your repo
4. Did you backup your repo?
5. edit the key values for the hash map `SEARCH_REPLACE_MAP` in kvmap.sh
6. test on the backup
7. *** touch .${0}.ready
8. run with real mapping changes
9. finalize changes test deployments, git push and whatever
10. git gc

*When DRY_RUN is 1, try to do no harm, print commands*


Use
---

After backing up your repository, editing the kvmap.sh you can test
with the `--dry-run` flag

```
    repo-text-replace --dry-run
```

After testing you can verify that you are ready by touching the ready
file and running with the `--exec` flag

```
    touch .repo-text-replace.ready
    repo-text-replace --exec
```
