import React, { Component } from 'react';
import { Editor as WYSIWYGEditor } from 'react-draft-wysiwyg';
import 'react-draft-wysiwyg/dist/react-draft-wysiwyg.css';
import './Editor.css';

class Editor extends Component {
  state = {
    contentState: "",
  }

  onContentStateChange: Function = (contentState) => {
    this.setState({
      contentState,
    });
  };

  render = () => {
    return(
      <WYSIWYGEditor
        toolbarClassName="toolbarClassName"
        wrapperClassName="wrapperClassName"
        editorClassName="editor-text-area"
        onContentStateChange={this.onContentStateChange}
      />
    )
  }
}

export default Editor;