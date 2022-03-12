import React, { useState } from 'react'

const UpdateStatus = () => {
  return (
    <div className='modalForm '>
      <h3>Update Status: </h3>
      <div>
        <label>Status: </label>
        <select>
          <option>Ongoing to Origin</option>
          <option>At Origin</option>
          <option>Ongoing to Destination</option>
          <option>At Destination</option>
          <option>Completed</option>
        </select>
      </div>
      <button className='modalButton'>Update</button>
      <button className='modalButton'>Cancel</button>
    </div>
  )
}

export default UpdateStatus
