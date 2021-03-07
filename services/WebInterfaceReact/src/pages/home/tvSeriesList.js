import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { Card, Typography, IconButton, Button, Input, TextField } from '@material-ui/core';
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
* TV Series list
* This component shows a list of tv series
*/

export class TVSeriesList extends Component {
  constructor(props) {
    super(props);
    this.handleChangePage = this.handleChangePage.bind(this);
  }

  handleChangePage(event, newPage) {
    const { loadNextPage } = this.props;
    loadNextPage(newPage + 1);
  };

  render() {
    const {
      tvSeries,
      page,
      totalPages,
      totalResults,
      loadSeasons,
    } = this.props;
    return (
      <div className="tvseries-list">
        <TableContainer>
           <Table stickyHeader>
              <TableHead>
                <TableRow>
                  <TableCell>Name</TableCell>
                  <TableCell>Original name</TableCell>
                  <TableCell />
                </TableRow>
              </TableHead>
              <TableBody>
                {tvSeries.map(tvSerie => (
                  <TableRow hover role="checkbox" key={tvSerie.id.toString()}>
                    <TableCell>
                      {tvSerie.name}
                    </TableCell>
                    <TableCell>
                      {tvSerie.original_name}
                    </TableCell>
                    <TableCell>
                      <Button variant="contained" onClick={() => loadSeasons(tvSerie.id)}>Select</Button>
                    </TableCell>
                  </TableRow>
                ))}
              </TableBody>
           </Table>
        </TableContainer>
        <TablePagination
          rowsPerPageOptions={[20]}
          component="div"
          count={totalResults}
          rowsPerPage={tvSeries.length}
          page={page-1}
          onChangePage={this.handleChangePage}
          onChangeRowsPerPage={() => {}}
      />
      </div>
    );
  }
}

TVSeriesList.propTypes = {
  history: PropTypes.objectOf(Object),
  tvSeries: PropTypes.arrayOf(Object),
  loadNextPage: PropTypes.func.isRequired,
  loadSeasons: PropTypes.func.isRequired,
};

const mapStateToProps = state => ({
  tvSeries: state.tvSeries,
  page: state.page,
  totalPages: state.totalPages,
  totalResults: state.totalResults,
});

export default connect(
  mapStateToProps,
)(TVSeriesList);
