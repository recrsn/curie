import React from 'react';
import './History.css';
import noData from '../../images/no-data.svg'

const History = () => (
    <div className="History">
        <div className="empty-state">
            <img height="250" src={noData} />
            <div>
                You don't have any URLs to show!
            </div>

        </div>
    </div>
)

export default History;
