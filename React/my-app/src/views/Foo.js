import React, {Component} from 'react'
import HelloWorld from '../components/HelloWorld'

class Foo extends React.Component {
    render() {
        return (
            <div className="hello">
                <HelloWorld msg="Welcome to Your React App"/>
            </div>
        )
    }

}

export default Foo
