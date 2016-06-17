import React, { Component } from 'react';
import getMuiTheme from 'material-ui/styles/getMuiTheme';
import {fade} from 'material-ui/utils/colorManipulator';
import LightBaseTheme from 'material-ui/styles/baseThemes/lightBaseTheme';


export default function withMaterialUI(ComposedComponent) {
  return class MaterialUI extends Component {
    static childContextTypes = {
      muiTheme: React.PropTypes.object
    }

    getChildContext() {
      return {
        muiTheme: getMuiTheme(LightBaseTheme)
      };
    }

    render() {
      const { context, ...other } = this.props;
      return <ComposedComponent {...other} />;
    }
  };
}