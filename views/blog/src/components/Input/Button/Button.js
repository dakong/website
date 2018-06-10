import React from 'react';
import PropTypes from 'prop-types';

import './Button.css';

const Button = (props) => (
  <button
    className="button-submit"
    onClick={() => props.onClick()}>{props.value}
  </button>
);

Button.propTypes = {
  value: PropTypes.string,
  onClick: PropTypes.func,
}
export default Button;