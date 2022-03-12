import React, { useState } from 'react'

const ShipperModal = ({ getShowModalStatus, getNewShipment }) => {
  const [addFormData, setAddFormData] = useState({
    origin: '',
    dest: '',
    loadingdate: '',
  })

  const handleChangeForm = (event) => {
    event.preventDefault()
    const fieldName = event.target.getAttribute('name')
    const fieldValue = event.target.value
    const newFormData = { ...addFormData }
    newFormData[fieldName] = fieldValue
    setAddFormData(newFormData)
  }

  const submitHandler = (event) => {
    event.preventDefault()
    if (!addFormData.origin || !addFormData.dest || !addFormData.loadingdate) {
      alert('please add origin, dest and loading date')
    } else {
      const newShipment = {
        shipment: 1,
        liscense: 'b 123',
        name: 'jen',
        origin: addFormData.origin,
        dest: addFormData.dest,
        loadingdate: addFormData.loadingdate,
        status: 'Active',
        action: '',
      }
      getNewShipment(newShipment)
    }
  }

  return (
    <div className='modalForm'>
      <form>
        <input
          type='text'
          name='origin'
          required='required'
          placeholder='enter origin'
          onChange={handleChangeForm}
        />
        <input
          type='text'
          name='dest'
          required='required'
          placeholder='enter dest'
          onChange={handleChangeForm}
        />
        <input
          type='text'
          name='loadingdate'
          required='required'
          placeholder='enter date'
          onChange={handleChangeForm}
        />
      </form>
      <button className='modalButton' onClick={submitHandler}>
        Submit
      </button>
      <button className='modalButton' onClick={() => getShowModalStatus()}>
        Cancel
      </button>
    </div>
  )
}

export default ShipperModal
