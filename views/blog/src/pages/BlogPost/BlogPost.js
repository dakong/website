import React, { Component } from 'react';
import Editor from '../../components/Editor';
import TextField from '../../components/Input/TextField';
import Button from '../../components/Input/Button';

import './BlogPost.css'

class BlogPost extends Component {
  state = {
    title: '',
  }
  onTextFieldChange = (value) => {
    this.setState({
      title: value,
    });
  }

  submitPost = () => {
    console.log('submit blog to the server');
  }

  render = () => {
    return (
      <div className="editor-container">
        <TextField
          onChange={this.onTextFieldChange}
          placeholder="Title"
          value={this.state.title}
        />
        <br />
        <Editor />
        <br />
        <Button
          onClick={this.submitPost}
          value="Post"
        />
      </div>
    );
  }
}

export default BlogPost