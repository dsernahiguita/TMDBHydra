const log = require('winston');
var json = require('../service.json');

const Configuration = require('./Configuration');

let config = null;

function initialize() {
  try {
    config = new Configuration(json);
  } catch (err) {
    throw err;
  }
}

function get(key) {
  if (!config) {
    throw new Error('Config is not yet initialized');
  }

  const value = config.get(key);
  if (!value) {
    log.warn(`Config parameter ${key} is not set. Please add it to the configuration file.`);
  }

  return value;
}

module.exports = {
  initialize,
  get,
};
