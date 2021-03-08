import React from 'react';
import PropTypes from 'prop-types';
import {ReactComponent as Logo} from '../../../assets/SVG/LogoTMDB.svg';
import './index.css';

export const Header = () => {
  return (
    <div className="header-layout">
      <div className="header">
        <Logo className='logo'/>
      </div>
    </div>
  );
};

Header.propTypes = {
  page: PropTypes.number,
  totalPages: PropTypes.number,
  onPressBack: PropTypes.func,
  children: PropTypes.oneOfType([
    PropTypes.arrayOf(PropTypes.node),
    PropTypes.node
  ]).isRequired
};

export default Header;
