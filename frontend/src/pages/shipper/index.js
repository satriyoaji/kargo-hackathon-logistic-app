import React, { useState } from 'react'
import ShipperModal from './shipperModal'
import UpdateStatus from './updateStatus'
import AllocateShipment from './allocateShipment'
import data from './mock-data.json'

import './style.scss'

const Shipper = () => {
  const [shipments, setShipments] = useState(data)
  const [showModal, setShowModal] = useState(false)
  const [selectedShipment, setSelectedShipment] = useState('')

  // status fields
  const [showStatus, setShowStatus] = useState(false)

  // allocate fields
  const [showAllocate, setShowAllocate] = useState(false)

  const getShowModalStatus = () => {
    setShowModal(!showModal)
  }

  const getNewShipment = (newShipment) => {
    setShipments([...shipments, newShipment])
    setShowModal(!showModal)
  }

  const getShowAllocateShipment = () => {
    setShowAllocate(!showAllocate)
  }

  const getShowStatus = () => {
    setShowStatus(!showStatus)
  }

  const handleUpdateClick = (shipment) => {
    setSelectedShipment(shipment.shipment)
    setShowStatus(true)
  }

  const handleAllocateClick = (shipment) => {
    setSelectedShipment(shipment.shipment)
    setShowAllocate(true)
  }

  return (
    <div className='body'>
      <button
        className='shipmentButton'
        onClick={() => setShowModal(!showModal)}
      >
        Add Shipment
      </button>
      {showModal && (
        <ShipperModal
          getShowModalStatus={getShowModalStatus}
          getNewShipment={getNewShipment}
        />
      )}
      <table>
        <thead>
          <tr>
            <th>Shipment</th>
            <th>Liscense</th>
            <th>Driver's Name</th>
            <th>Origin</th>
            <th>Destination</th>
            <th>Loading Date</th>
            <th>Status</th>
            <th>Action</th>
          </tr>
        </thead>
        <tbody>
          {shipments.map((shipment) => (
            <tr>
              <td>{shipment.shipment}</td>
              <td>{shipment.liscense}</td>
              <td>{shipment.name}</td>
              <td>{shipment.origin}</td>
              <td>{shipment.dest}</td>
              <td>{shipment.loadingdate}</td>
              <td>{shipment.status}</td>
              <td>
                <button onClick={() => handleUpdateClick(shipment)}>
                  Update Status
                </button>
                <button onClick={() => handleAllocateClick(shipment)}>
                  Allocate Shipment
                </button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
      {showStatus && (
        <UpdateStatus id={selectedShipment} getShowStatus={getShowStatus} />
      )}
      {showAllocate && (
        <AllocateShipment
          id={selectedShipment}
          getShowAllocateShipment={getShowAllocateShipment}
        />
      )}
    </div>
  )
}

export default Shipper
