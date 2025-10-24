const fs = require('fs');
let json = JSON.parse(fs.readFileSync('angular.json', 'utf8'));
json.projects['shell-app'].architect.serve = {
  'builder': '@angular-architects/native-federation:serve',
  'options': { 'port': 4200, 'publicHost': 'http://localhost:4200' }
};
json.projects['shell-app'].architect.build = {
  'builder': '@angular-architects/native-federation:build',
  'options': { 'publicHost': 'http://localhost:4200' }
};
fs.writeFileSync('angular.json', JSON.stringify(json, null, 2));
console.log('Updated angular.json successfully!');