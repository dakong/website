import React from 'react';
import  PropTypes from 'prop-types';

import './TextField.css';

const TextField = (props) => {
  return (
    <div className="text-field-container">
      { props.value !== "" &&
        <label className="text-field-label">
          {props.placeholder}
        </label>
      }
      <br/>
      <input
        className="text-field-input"
        placeholder={props.placeholder}
        value={props.value}
        onChange={(e) => {props.onChange(e.target.value)}}
      />
    </div>
  );
}

TextField.propTypes = {
  value: PropTypes.string,
  placeholder: PropTypes.string,
  onChange: PropTypes.func,
}

export default TextField;