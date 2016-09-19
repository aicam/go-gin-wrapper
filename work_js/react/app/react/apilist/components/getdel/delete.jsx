import React from 'react'
import Options from './options.jsx'

export default class Delete extends React.Component {
  //like getInitialState()
  constructor(props) {
    super(props)

    this.clickBtnEvt = this.clickBtnEvt.bind(this)
  }

  //Click delete button
  clickBtnEvt(e) {
    console.log("[Get]:clickBtnEvt(e)")
    let idx = this.refs.slctDelIDs.selectedIndex
    let val = this.refs.slctDelIDs.options[idx].text

    //call event for post btn click
    this.props.btn.call(this, val)
  }

  render() {
    let options = this.props.ids.map(function (id) {
      let key='del_'+id
      return (
        <Options key={key} id={id} />
      )
    })

    return (
        <div className="col-md-6 col-sm-6 col-xs-12">
          <div className="panel panel-danger">
            <div className="panel-heading">DELETE API</div>
            <div className="panel-body">
                <div className="form-group">
                  <label>Select user id</label>
                  <select id="slctDelIds" className="form-control" ref="slctDelIDs">
                    {options}
                  </select>
                </div>
                <button id="delBtn" type="button" onClick={this.clickBtnEvt} className="btn btn-danger">Delete User</button>
            </div>
          </div>
        </div>
    )
  }
}
