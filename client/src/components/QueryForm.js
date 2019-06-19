import React, { Component } from 'react';

const QUERY_API_URI = 'http://localhost:5000/api/query';

class QueryForm extends Component {
  constructor(props) {
    super(props);
    this.state = { Id: '', };
  }

  handleSubmit = (event) => {
    event.preventDefault();

    const { Id } = this.state;
    const fetchOpts = {
      method: 'get',
      mode: 'cors',
      headers: {
        "Content-Type": "application/json",
      },
      qs: { studId: Id },
    };

    fetch(QUERY_API_URI, fetchOpts)
      .then((data) => console.log('Request succeeded with JSON response', data))
      .catch((error) => console.error('Request failed', error));
  }

  handleChange = (event) => {
    const { name, value } = event.target;
    this.setState({ [name]: value });
  }

  render() {
    return (
      <div>
        <h1 className="display-4">
          Query
        </h1>
        <form onSubmit={this.handleSubmit}>
          <div className="form-group">
            <label htmlFor="exampleInputID2">
              ID студента
            </label>
            <input
              type="text"
              name="Id"
              value={this.state.Id}
              onChange={this.handleChange}
              className="form-control"
              id="exampleInputID2"
              placeholder="Введите идентификатор студента"
            />
          </div>
          <button type="submit" className="btn btn-success">
            Отправить
          </button>
        </form>
      </div>
    );
  }
}

export default QueryForm;