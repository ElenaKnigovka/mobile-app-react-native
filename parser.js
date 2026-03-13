const fs = require('fs');
const path = require('path');

class Parser {
  constructor(filePath) {
    this.filePath = filePath;
    this.data = {};
  }

  readFileSync() {
    try {
      const content = fs.readFileSync(this.filePath, 'utf8');
      return content;
    } catch (error) {
      throw new Error(`Error reading file: ${error}`);
    }
  }

  parseFile() {
    const content = this.readFileSync();
    const lines = content.split('\n');
    lines.forEach((line) => {
      const [key, value] = line.split('=');
      this.data[key.trim()] = value.trim();
    });
    return this.data;
  }

  saveData(data) {
    try {
      fs.writeFileSync(this.filePath, '');
      Object.keys(data).forEach((key) => {
        fs.appendFileSync(this.filePath, `${key}=${data[key]}\n`);
      });
    } catch (error) {
      throw new Error(`Error writing to file: ${error}`);
    }
  }
}

module.exports = Parser;