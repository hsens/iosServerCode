class App extends React.Component {
  render(){
    return (<Home />);
  }
}
class Home extends React.Component{

    constructor(props) {
      super(props);
      this.authenticate = this.authenticate.bind(this);
    }
    authenticate() {
      
      ReactDOM.render(<Main/>, document.getElementById('app'));
      ReactDOM.render(<Nav />, document.getElementById('main_nav'));
    }
  
      render(){
          return(
          <div className="container">
          <div className="col-xs-8 col-xs-offset-2 jumbotron text-center">
            <h1><b>h</b>sens</h1>
            <p>ios Dashboard</p>
            <p>Sign in to get access </p>
            <a onClick={this.authenticate} className="btn btn-primary btn-lg btn-login btn-block">Sign In</a>
            
          </div>
        </div>
          )
      }
      
  }

class Main extends React.Component{
  render() {
    return (
       <div>
          <Header/>
          <Content/>
          
       </div>
    );
 }
}
class Header extends React.Component {
  render() {
     return (
        <div>
           <h1>Header</h1>
        </div>
     );
  }
}
class Content extends React.Component {
  getall(e){
    fetch("https://ios.pandawalab.id:3000/api/people", {
    "method": "GET",
    "headers": {
      "Authorization": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.tcAmnNdj5xSgsEGfLg-SXsYuP1-Efllbh5Y9k5u7NTw"
  }
  
})
.then(response => response.json())
.then(response => {
  this.setState({
    friends: response
  })
  console.log(response)
})
.catch(err => { console.log(err); 
});
  }
  

  render() {
     return (
        <div>
           <h2>Content</h2>
           <p>The content text!!!</p>
           <img className="map.jpg" src="./images" alt="" />
           <button className="btn btn-info" type='button' onClick={(e) => this.getall(e)}>
                    Update
                </button>
                <div id="result"></div>
        </div>
       
     );
  }
}
class Nav extends React.Component {
  render() {
     return (
      <ul>
      <li><a href="#">Dashboard</a></li>
      <li><a href="#">Active List</a></li>
      <li><a href="#">Inactive List</a></li>
      <li><a href="#">All Messages</a></li>
      </ul>
     );
  }
}
class Friends extends React.Component {
  render() {
      return (
          <table>
              <thead>
                  <tr>
                      <th>ID</th>
                      <th>Name</th>
                      <th>Since</th>
                  </tr>
              </thead>
              <tbody>
                  {this.props.friends && this.props.friends.map(friend => {
                      return <tr>
                          <td>{friend.imei}</td>
                          <td>{friend.first_name}</td>
                          <td>{friend.since}</td>
                      </tr>
                  })}
              </tbody>
          </table>
      );
  }
}


ReactDOM.render(<App />, document.getElementById('app'));