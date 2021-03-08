import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { Button } from '@material-ui/core';
import SearchIcon from '@material-ui/icons/Search';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TablePagination from '@material-ui/core/TablePagination';
import TableRow from '@material-ui/core/TableRow';
import TableSortLabel from '@material-ui/core/TableSortLabel';
import './index.css';

/**
* Seasons
* This component shows a list of seasons
*/

export class Seasons extends Component {
  constructor(props) {
    super(props);
  }

  render() {
    const {
      seasons,
      nameTVSerie,
      tvSerieId,
      loadEpisodes,
    } = this.props;
    return (
      <div className="tvseries-list">
        <TableContainer>
           <Table stickyHeader>
             <TableHead>
               <TableRow>
                 <TableCell
                  colSpan={3}
                  align="center"
                 >
                    {nameTVSerie} - Seasons
                  </TableCell>
                </TableRow>
             </TableHead>
              <TableHead>
                <TableRow>
                  <TableCell>Name</TableCell>
                  <TableCell>Overview</TableCell>
                  <TableCell />
                </TableRow>
              </TableHead>
              <TableBody>
                {seasons.map(season => (
                  <TableRow hover role="checkbox" key={season.id.toString()}>
                    <TableCell>
                      {season.name}
                    </TableCell>
                    <TableCell>
                      {season.overview}
                    </TableCell>
                    <TableCell>
                      <Button variant="contained" onClick={() => loadEpisodes(tvSerieId, season.season_number)}>Select</Button>
                    </TableCell>
                  </TableRow>
                ))}
              </TableBody>
           </Table>
        </TableContainer>
      </div>
    );
  }
}

Seasons.propTypes = {
  history: PropTypes.objectOf(Object),
  seasons: PropTypes.arrayOf(Object),
  nameTVSerie: PropTypes.string.isRequired,
  tvSerieId: PropTypes.number.isRequired,
  loadEpisodes: PropTypes.func.isRequired,
};

const mapStateToProps = state => ({
  seasons: state.seasons,
  nameTVSerie: state.nameTVSerie,
  tvSerieId: state.tvSerieId,
});

export default connect(
  mapStateToProps,
)(Seasons);
