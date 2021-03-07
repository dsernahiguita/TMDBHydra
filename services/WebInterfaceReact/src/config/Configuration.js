/**
* Configuration
*/
class Configuration {
  constructor(params) {
    this.params = params;
  }

  get(key) {
    return this.params[key];
  }
}

module.exports = Configuration;
