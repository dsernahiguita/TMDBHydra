import AxiosClient from './axiosClient';

const ENDPOINT_AUTH = ":4060/api/frontend";

/**
* TMDB Hydra client
* Implements the api calls to talk with the service HydraClient
*/
export default class TMDBHydraClient {
  constructor() {
    this.client = new AxiosClient();
  }

  /**
  * Get tv series
  * @param string tvSerie
  * @param int page
  */
  async getTVSeries(tvSerie, page) {
    const endPoint = `${this.client.getURLWithPrefix(ENDPOINT_AUTH)}/tvserie?query=${tvSerie}&page=${page}`;
    const response = await this.client.get(endPoint);
    return response;
  }

  /**
  * Get seasons
  * @param int tvSerieId
  */
  async getSeasons(tvSerieId) {
    const endPoint = `${this.client.getURLWithPrefix(ENDPOINT_AUTH)}/seasons?tvSerieId=${tvSerieId}`;
    const response = await this.client.get(endPoint);
    return response;
  }

  /**
  * Get episodes
  * @param int tvSerieId
  * @param int season
  */
  async getEpisodes(tvSerieId, season) {
    const endPoint = `${this.client.getURLWithPrefix(ENDPOINT_AUTH)}/episodes?tvSerieId=${tvSerieId}&season=${season}`;
    const response = await this.client.get(endPoint);
    return response;
  }
}
