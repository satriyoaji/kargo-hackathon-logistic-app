import React, { useState } from 'react'

const AllocateShipment = ({ id, getShowAllocateShipment }) => {
  const handleCancel = (event) => {
    event.preventDefault()
    getShowAllocateShipment(false)
  }

  const handleAllocate = (event) => {
    event.preventDefault()
    // api
    getShowAllocateShipment(false)
  }
  return (
    <div className='modalForm '>
      <h3>Allocate Shipments: {id}</h3>
      <div>
        <label>Truck: </label>
        <select></select>
      </div>
      <div>
        <label>Driver: </label>
        <select></select>
      </div>
      <button onClick={handleAllocate} className='modalButton'>
        Allocate{' '}
      </button>
      <button onClick={handleCancel} className='modalButton'>
        Cancel
      </button>
    </div>
  )
}

export default AllocateShipment
