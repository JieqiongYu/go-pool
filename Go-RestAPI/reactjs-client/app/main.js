'use strict';
const React = require('react');
const ReactDom = require('react-dom')
import EmployeeApp from './components/employee-app.jsx'

ReactDom.render(
    <EmployeeApp />,
    document.getElementById('react')
)