import axios from 'axios';

/**
* Axios Client
*/
export default class AxiosClient {
  constructor() {
    this.backendHost = window.location.hostname;
    this.backendProtocol = window.location.protocol;
  }

  /**
  * Get
  * call the methode get of axios
  * @param string endPoint
  * @param objects params
  * @return object response.data
  */
  async get(endPoint, params, config = null) {
    try {
      const response = await axios.get(
        endPoint,
        { data: params },
        !config ? this.configHeaders() : config,
      );
      return response.data
    } catch (error) {
      const errorMessage = {
        message: error.Message,
        errorId: error.ErrorId,
      };
      throw(errorMessage);
    }
  }

  /**
  * Default error handling
  * @param object error
  */
  defaultErrorHandling(error) {
    const errorMessage = {
      message: error.Message,
      errorId: error.ErrorId,
    };
    throw(errorMessage);
  }

  /**
  * Post
  * call the methode post of axios
  * @param string endPoint
  * @param object params
  * @param object config
  * @return object response.data
  */
  async post(endPoint, params, config = null, errorHandling = this.defaultErrorHandling) {
    try {
      const response = await axios.post(
        endPoint,
        params,
        !config ? this.configHeaders : config,
      );
      return response.data
    } catch (error) {
      errorHandling(error);
    }
  }

  /**
  * put
  * call the methode post of axios
  * @param string endPoint
  * @param object params
  * @param object config
  * @return object response.data
  */
  async put(endPoint, params, config = null, errorHandling = this.defaultErrorHandling) {
    try {
      const response = await axios.post(
        endPoint,
        params,
        !config ? this.configHeaders : config,
      );
      return response.data
    } catch (error) {
      errorHandling(error);
    }
  }

  /**
  * Get backend host url
  * @return string url
  */
  getBackendHostURL() {
    return `${this.backendProtocol}//${this.backendHost}`
  }

  /**
  * Get url with prefix
  * @param string prefix
  * @return string url
  */
  getURLWithPrefix(prefix) {
    return this.getBackendHostURL() + prefix;
  }
  
  /**
  * Config header
  */
  configHeaders() {
    return {
      crossdomain: true,
      headers: {
        'content-type': 'text/json',
        'Access-Control-Allow-Origin': '*',
        'Access-Control-Allow-Headers': 'Content-Type, Authorization',
      }
    };
  }
}
