import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { Card, Typography, Button } from '@material-ui/core';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import CardHeader from '@material-ui/core/CardHeader';
import CardMedia from '@material-ui/core/CardMedia';
import CardContent from '@material-ui/core/CardContent';
import Avatar from '@material-ui/core/Avatar';
import './index.css';

/**
* Episodes
* This component shows a list of episodes
*/

export class Episodes extends Component {
  constructor(props) {
    super(props);
    this.selectEpisode = this.selectEpisode.bind(this);
    this.state = {
      episode: null,
    }
  }

  /**
  * Select episode
  * @param object episode
  */
  selectEpisode(episode) {
    this.setState({ episode });
  }


  render() {
    const {
      episodes,
      nameTVSerie,
      nameSeason,
    } = this.props;
    const { episode } = this.state;
    if (episode) {
      return (
        <div className="episode">
          <Card>
            <CardHeader
              avatar={
                <Avatar aria-label="recipe">
                  R
                </Avatar>
              }
              title={episode.name}
              subheader={`TV Serie: ${nameTVSerie} Season: ${nameSeason}`}
            />
            <CardMedia
              image="/static/images/cards/paella.jpg"
              title="Paella dish"
            />
            <CardContent>
              <Typography variant="body2" color="textSecondary" component="p">
                {episode.overview}
              </Typography>
            </CardContent>
          </Card>
        </div>
      );
    }

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
                    {nameTVSerie} {nameSeason} - Episodes
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
                {episodes.map(episode => (
                  <TableRow hover role="checkbox" key={episode.id.toString()}>
                    <TableCell>
                      {episode.name}
                    </TableCell>
                    <TableCell>
                      {episode.overview}
                    </TableCell>
                    <TableCell>
                      <Button variant="contained" onClick={() => this.selectEpisode(episode)}>Select</Button>
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

Episodes.propTypes = {
  history: PropTypes.objectOf(Object),
  episodes: PropTypes.arrayOf(Object),
  nameTVSerie: PropTypes.string.isRequired,
  nameSeason: PropTypes.string.isRequired,
};

const mapStateToProps = state => ({
  episodes: state.episodes,
  nameTVSerie: state.nameTVSerie,
  nameSeason: state.nameSeason,
});

export default connect(
  mapStateToProps,
)(Episodes);
