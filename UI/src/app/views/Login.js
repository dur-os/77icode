import React, { PropTypes, Component }	from 'react';
import { connect }	from 'react-redux';
import { Paper, TextField, RaisedButton }	from 'material-ui';
import { loginUser }	from '../actions/user';


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
    let isValidate = false;
    if (!userName) {
      this.refs.userName.setState(Object.assign({},
         this.refs.userName.state, { errorText: '用户名不能为空' }));
      isValidate = true;
    } else {
      this.refs.userName.setState(Object.assign({},
         this.refs.userName.state, { errorText: '' }));
    }

    const password = this.refs.password.state.hasValue;
    if (!password) {
      this.refs.password.setState(Object.assign({},
         this.refs.password.state, { errorText: '密码不能为空' }));
      isValidate = true;
    } else {
      this.refs.password.setState(Object.assign({},
         this.refs.password.state, { errorText: '' }));
    }
    if (isValidate) {
      return;
    }
    this.props.loginUser(this.refs.userName.getValue(), this.refs.password.getValue());
  }

  renderErrorMessage() {
    const { errorMessage } = this.props;
    console.log('errr', errorMessage);
    if (!errorMessage) {
      return null;
    }

    return (
      <p style={ { backgroundColor: '#e99', padding: 10 } }>
        <b>{ errorMessage }</b>
      </p>
    );
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
          { this.renderErrorMessage() }
        </Paper>
      </div>
    );
  }
}

Login.propTypes = {
  errorMessage: PropTypes.string,
  loginUser: React.PropTypes.func.isRequired
};

export default connect((state, ownProps) => {
  const { user } = state;
  const userName = ownProps.params.name;
  return {
    errorMessage: state.errorMessage,
    user,
    userName
  };
}, { loginUser })(Login);
