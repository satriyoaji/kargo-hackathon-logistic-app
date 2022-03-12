import React, { useState } from 'react'

const AllocateShipment = () => {
  return (
    <div className='modalForm '>
      <h3>Allocate Shipments: </h3>
      <div>
        <label>Truck: </label>
        <select></select>
      </div>
      <div>
        <label>Driver: </label>
        <select></select>
      </div>
      <button className='modalButton'>Allocate </button>
      <button className='modalButton'>Cancel</button>
    </div>
  )
}

export default AllocateShipment
