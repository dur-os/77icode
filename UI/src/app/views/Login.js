import React, { PropTypes, Component } from 'react';
import { connect } from 'react-redux';
import { Paper, TextField, RaisedButton } from 'material-ui';

class Login extends Component {
  static contextTypes = {
    muiTheme: PropTypes.object.isRequired
  }

  getStyles() {
    return {
      text: {
        width: '100%'
      },
      submit: {
        marginTop: 10,
        marginBottom: 20,
        height: 50,
        width: '100%'
      },
      submitLabel: {
        textTransform: 'none',
        fontSize: 20,
        letterSpacing: 1,
        fontFamily: 'Monaco,"Courier New",sans-serif'
      }
    };
  }

  submit(event) {
    if (event.type === 'keydown' && event.keyCode !== 13) return;

    // const { dispatch } = this.props;
    // const actions = bindActionCreators(AuthActions, dispatch);

    const userName = this.refs.userName.state.hasValue;
    const password = this.refs.password.state.hasValue;
    console.log(userName, password);
  }

  render() {
    const styles = this.getStyles();
    return (
      <div className="login">
        <Paper className="paper">
          <div className="loginTitle">Sign In</div>
          <TextField
            ref="userName"
            style={ styles.text }
            hintText="UserName"
            floatingLabelText="UserName"
            onKeyDown={ ::this.submit }
          /><br />
          <TextField
            ref="password"
            style={ styles.text }
            hintText="Password"
            floatingLabelText="Password"
            onKeyDown={ ::this.submit }
            type="password"
          /><br />
          <RaisedButton
            labelStyle={ styles.submitLabel }
            style={ styles.submit }
            label="Sign In"
            onTouchTap={ ::this.submit }
            secondary
          />
        </Paper>
      </div>
    );
  }
}


export default connect(state => ({ user: state }))(Login);
