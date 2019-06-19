import React, { Component } from 'react';

const INVOKE_API_URI = 'http://localhost:5000/api/invoke';

class InvokeForm extends Component {
  constructor(props) {
    super(props);
    this.state = {
      studId: '',
      studFirstName: '',
      studLastName: '',
      studMiddleName: '',
      studPlaceOfBirth: '',
      studDateOfBirth: '',
      studPassNum: '',
      studGender: '',
      studMaritalStatus: '',
    };
  }

  handleSubmit = (event) => {
    event.preventDefault();

    const {
      studId,
      studFirstName,
      studLastName,
      studMiddleName,
      studPlaceOfBirth,
      studDateOfBirth,
      studPassNum,
      studGender,
      studMaritalStatus,
    } = this.state;
    const body = JSON.stringify({
      studId,
      studFirstName,
      studLastName,
      studMiddleName,
      studPlaceOfBirth,
      studDateOfBirth,
      studPassNum,
      studGender,
      studMaritalStatus,
    });
    const fetchOpts = {
      method: 'post',
      mode: 'cors',
      headers: {
        "Content-Type": "application/json",
      },
      body,
    };

    fetch(INVOKE_API_URI, fetchOpts)
      .then((data) => console.log('Request succeeded with JSON response', data))
      .catch((error) => console.error('Request failed', error));
  }

  handleChange = (event) => {
    const { name, value } = event.currentTarget;
    this.setState({ [name]: value });
  }

  render() {
    return (
      <div>
        <h1 className="display-4">Invoke</h1>
        <form onSubmit={this.handleSubmit}>
          <div className="form-group">
            <label htmlFor="idStud">
              ID
            </label>
            <input
              type="text"
              className="form-control"
              value={this.state.studId}
              onChange={this.handleChange}
              name="studId"
              id="idStud"
              placeholder="Введите ID"
            />
          </div>
          <div className="form-group">
            <label htmlFor="studName">
              Имя
            </label>
            <input
              type="text"
              className="form-control"
              value={this.state.studFirstName}
              onChange={this.handleChange}
              name="studFirstName"
              id="studName"
              placeholder="Введите имя"
            />
          </div>
          <div className="form-group">
            <label htmlFor="studLName">
              Фамилия
            </label>
            <input
              type="text"
              className="form-control"
              value={this.state.studLastName}
              onChange={this.handleChange}
              name="studLastName"
              id="studLName"
              placeholder="Введите фамилию"
            />
          </div>
          <div className="form-group">
            <label htmlFor="studMName">
              Отчество
            </label>
            <input
              type="text"
              className="form-control"
              value={this.state.studMiddleName}
              onChange={this.handleChange}
              name="studMiddleName"
              id="studMName"
              placeholder="Введите отчество"
            />
          </div>
          <div className="form-group">
            <label htmlFor="placeOfBirth">
              Место рождения
            </label>
            <input
              type="text"
              className="form-control"
              value={this.state.studPlaceOfBirth}
              onChange={this.handleChange}
              name="studPlaceOfBirth"
              id="placeOfBirth"
              placeholder="Введите место рождения"
            />
          </div>
          <div className="form-group">
            <label htmlFor="studDOB">
              Дата рождения
            </label>
            <input
              type="date"
              className="form-control"
              value={this.state.studDateOfBirth}
              onChange={this.handleChange}
              name="studDateOfBirth"
              id="studDOB"
              placeholder="Введите дату рождения"
            />
          </div>
          <div className="form-group">
            <label htmlFor="studPN">
              Серия, номер паспорта
            </label>
            <input
              type="text"
              className="form-control"
              value={this.state.studPassNum}
              onChange={this.handleChange}
              name="studPassNum"
              id="studPN"
              placeholder="Введите серию, номер паспорта"
            />
          </div>
          <div className="form-group mb-3">
            <div className="input-group-prepend">
              <label htmlFor="studGN">
                Пол
              </label>
            </div>
            <select
              type="text"
              className="form-control"
              value={this.state.studGender}
              onChange={this.handleChange}
              name="studGender"
              id="studGN"
            >
              <option value="">
                Выберите...
              </option>
              <option value="М">
                М
              </option>
              <option value="Ж">
                Ж
              </option>
            </select>
          </div>
          <div className="form-group mb-3">
            <div className="input-group-prepend">
              <label htmlFor="idStud">
                Семейное положение
              </label>
            </div>
            <select
              type="text"
              className="form-control"
              value={this.state.studMaritalStatus}
              onChange={this.handleChange}
              name="studMaritalStatus"
              id="studMStatus"
            >
              <option value="">
                Выберите...
              </option>
              <option value="Холост">
                Холост
              </option>
              <option value="Женат">
                Женат
              </option>
            </select>
          </div>
          <button type="submit" className="btn btn-success">
            Отправить
          </button>
        </form>
      </div>
    );
  }
}

export default InvokeForm;