import {
  GET_TVSERIES,
  GET_SEASONS,
  GET_EPISODES,
} from '../actions/types';

const initialState = {
  page: -1,
  totalPages: -1,
  totalResults: -1,
  tvSeries: [],
  tvSerieId: -1,
  seasons: [],
  episodes:[],
  nameTVSerie: '',
  numberEpisodes: -1,
  numberSeasons: -1,
  nameSeason: '',
};

function rootReducer(state = initialState, action) {
  const { type, payload } = action;
  switch (type) {
    case GET_TVSERIES:
      return {
        ...state,
        tvSeries: payload.tvSeries,
        page: payload.page,
        totalPages: payload.totalPages,
        totalResults: payload.totalResults,
        episodes: [],
        seasons: [],
        nameTVSerie: '',
        numberEpisodes: -1,
        numberSeasons: -1,
      };
    case GET_SEASONS:
      return {
        ...state,
        tvSerieId: payload.tvSerieId,
        seasons: payload.seasons,
        nameTVSerie: payload.nameTVSerie,
        numberEpisodes: payload.numberEpisodes,
        numberSeasons: payload.numberSeasons,
        tvSeries: [],
      };
    case GET_EPISODES:
      return {
        ...state,
        episodes: payload.episodes,
        nameSeason: payload.nameSeason,
        seasons: [],
      };
    default:
      break;
  }
  return state;
}

export default rootReducer;
