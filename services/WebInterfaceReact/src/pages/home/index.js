import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { Button, TextField, IconButton, CircularProgress } from '@material-ui/core';
import SearchIcon from '@material-ui/icons/Search';
import TVSeriesList from './tvSeriesList';
import Seasons from './seasons';
import Episodes from './episodes';
import { getTVSeries, getSeasons, getEpisodes } from '../../store/actions';
import './index.css';

/**
* Home
* this page allows the user searches for a tvserie, select a season and select
* a episode
*/
export class Home extends Component {
  constructor(props) {
    super(props);
    this.loadData = this.loadData.bind(this);
    this.loadNextPage = this.loadNextPage.bind(this);
    this.loadSeasons = this.loadSeasons.bind(this);
    this.loadEpisodes = this.loadEpisodes.bind(this);
    this.handleTextFieldChange = this.handleTextFieldChange.bind(this);
    this.state = {
      tvSerie: '',
      loading: false,
    }
  }

  /**
  * Load data
  * This function call the service getTVSeries from the backend
  */
  async loadData() {
    try {
      const { tvSerie } = this.state;
      this.setState({ loading: true });
      await this.props.getTVSeries(tvSerie);
      this.setState({ loading: false });
    } catch (error) {
      this.setState({ loading: false });
      console.error(error);
    }
  }

  /**
  * Load next page
  * @param number page
  * This function call the service getTVSeries from the backend
  */
  async loadNextPage(page) {
    try {
      const { tvSerie } = this.state;
      this.setState({ loading: true });
      await this.props.getTVSeries(tvSerie, page);
      this.setState({ loading: false });
    } catch (error) {
      this.setState({ loading: true });
      console.error(error);
    }
  }

  /**
  * Load seasons
  * @param number tvSerieId
  * This function call the service getSeasons from the backend
  */
  async loadSeasons(tvSerieId) {
    try {
      this.setState({ loading: true });
      await this.props.getSeasons(tvSerieId);
      this.setState({ loading: false });
    } catch (error) {
      this.setState({ loading: false });
      console.error(error);
    }
  }

  /**
  * Load episodes1
  * @param number tvSerieId
  * @param number season
  * This function call the service getEpisodes from the backend
  */
  async loadEpisodes(tvSerieId, season) {
    try {
      this.setState({ loading: true });
      await this.props.getEpisodes(tvSerieId, season);
      this.setState({ loading: false });
    } catch (error) {
      this.setState({ loading: false });
      console.error(error);
    }
  }

  /**
  * Handle text field change
  * @param object e: value = e.target
  */
  handleTextFieldChange(e) {
    this.setState({ tvSerie: e.target.value });
  }

  render() {
    const {
      tvSeries,
      seasons,
      episodes,
    } = this.props;
    const { tvSerie, loading } = this.state;
    return (
      <div className="home-layout">
        <div className="title-layout">
          <div className="title">
            TVSeries, Seasons and Episodes
          </div>
          <div className="subtitle">
            By HYDRA
          </div>
        </div>
        <div className="body">
          <div className="input-field">
            <TextField
              id="tvSerie"
              label="Please enter the name of a TV Serie"
              variant="outlined"
              value={tvSerie}
              onChange={this.handleTextFieldChange}
              fullWidth
            />
          </div>
          <div className="button">
            <IconButton
              color="primary"
              aria-label="Search"
              onClick={this.loadData}
            >
              <SearchIcon fontSize="medium"/>
            </IconButton>
          </div>
        </div>
        {loading && (
          <div className="loading">
            <CircularProgress />
          </div>
        )}
        {episodes.length > 0 && (
          <Episodes

          />
        )}
        {seasons.length > 0 && (
          <Seasons
            loadEpisodes={this.loadEpisodes}
          />
        )}

        {tvSeries.length > 0 && (
          <TVSeriesList
            loadNextPage={this.loadNextPage}
            loadSeasons={this.loadSeasons}
          />
        )}
      </div>
    );
  }
}

Home.propTypes = {
  history: PropTypes.objectOf(Object),
  getTVSeries: PropTypes.func.isRequired,
  getSeasons: PropTypes.func.isRequired,
  tvSeries: PropTypes.arrayOf(Object),
  seasons: PropTypes.arrayOf(Object),
  episodes: PropTypes.arrayOf(Object),
  page: PropTypes.number.isRequired,
};

const mapStateToProps = state => ({
  tvSeries: state.tvSeries,
  seasons: state.seasons,
  page: state.page,
  episodes: state.episodes
});

export default connect(
  mapStateToProps,
  { getTVSeries, getSeasons, getEpisodes }
)(Home);
