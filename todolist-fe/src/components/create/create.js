import React, { useState } from "react";
import { Button, Form, Checkbox} from 'semantic-ui-react';
import axios from "axios";
import { useNavigate } from 'react-router-dom';

export default function Create() {
    let navigate = useNavigate();
    const [name, setTodoName] = useState('');
    const [checked, setChecked] = useState(false);

    const post = () => {
        axios.post("http://localhost:8000/todolist", {
            name, 
            status:checked
        }).then(() => {
            navigate('/')
        })
    }

    console.log(checked)

    return (
        <div>
            <Form>
                <Form.Field>
                    <label>Todo Name</label>
                    <input name="name" placeholder='Todo Name' onChange={(e) => setTodoName(e.target.value)} />
                </Form.Field>
                <Form.Field>
                    <Checkbox label="Check for completed todo" onChange={(e) => setChecked(!checked)}></Checkbox>
                </Form.Field>
                <Button type='submit' onClick={post}>Submit</Button>
            </Form>
        </div>
    )
}