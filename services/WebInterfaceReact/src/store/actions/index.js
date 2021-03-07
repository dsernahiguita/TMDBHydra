import { tmdbHydraClient } from '../../api';
import {
  GET_TVSERIES,
  GET_SEASONS,
  GET_EPISODES,
} from './types';


/**
* Get tv series
* @param string query
* @param int page
*/
export const getTVSeries = (query, page) => async (dispatch) => {
  const tvSeries = await tmdbHydraClient.getTVSeries(query, page);

  dispatch({
    type: GET_TVSERIES,
    payload: {
      page: tvSeries.page,
      totalPages: tvSeries.total_pages,
      totalResults: tvSeries.total_results,
      tvSeries: tvSeries.results,
    },
  });
};


/**
* Get Seasons
* @param int tvSerieId
*/
export const getSeasons = (tvSerieId) => async (dispatch) => {
  const seasons = await tmdbHydraClient.getSeasons(tvSerieId);
  dispatch({
    type: GET_SEASONS,
    payload: {
      tvSerieId,
      seasons: seasons.seasons,
      nameTVSerie: seasons.name,
      numberEpisodes: seasons.number_of_episodes,
      numberSeasons: seasons.number_of_seasons,
    },
  });
};

/**
* Get episodes
* @param int tvSerieId
* @param int season
*/
export const getEpisodes = (tvSerieId, season) => async (dispatch) => {
  const episodes = await tmdbHydraClient.getEpisodes(tvSerieId, season);
  dispatch({
    type: GET_EPISODES,
    payload: {
      episodes: episodes.episodes,
      nameSeason: episodes.name,
    }
  });
};
