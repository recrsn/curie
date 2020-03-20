import React from 'react';
import URIMaker from '../URIMaker/URIMaker';
import History from '../history/History';

import './Dashboard.css';

const Dashboard = () => (
    <div className="Dashboard">
        <section className="section-uri-maker">
            <div className="container">
                <URIMaker />
            </div>
        </section>
        <section className="section-history">
            <div className="container">
                <History />
            </div>
        </section>
    </div>
)

export default Dashboard;
