To replicate TournyMasterBot's environment, follow the steps below. If you already have a database server set up, you can simply swap in your information where appropriate to use that instead.

#Assumptions: 
- Mysql port is 3306, plan accordingly.

#Non Required Tools
- Mysql Workbench: https://dev.mysql.com/downloads/workbench/

#Setup
1. Install Mysql Server
2. Allow Mysql through the firewall
3. Open the command prompt to where you've installed mysql, by default for mysql 8 this is located at '' and log in through the cli
4. Execute the following commands:

```sql
CREATE USER 'root'@'%' IDENTIFIED BY '{password}';
GRANT ALL PRIVILEGES ON *.* TO 'root'@'%';
```
