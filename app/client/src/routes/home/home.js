import React from 'react'
import { gql } from 'apollo-boost';
import { Query, Mutation, graphql } from 'react-apollo';
import DeleteForeverIcon from '@material-ui/icons/DeleteForever';

const GET_NOT_TODO = gql`
	query GetNotTodos{
		getNotTodos {
			id
			name
			description
			complete
		}
	}
`

const CREATE_NOT_TODO = gql`
	mutation CreateNotTodo($name: String, $description: String) {
		createNotTodo(name: $name, description: $description) {
			name
			description
			complete
		}
	}
`

const DELETE_TODO = gql`
	mutation DeleteTodo($id: String) {
		deleteTodo(id: $id) {
			id
		}
	}
`

class Home extends React.Component {
	constructor(props){
		super(props)

		this.state = {
			name:'',
			description: ''
		}
	}

	onChange = event => {
		this.setState({ [event.target.name]: event.target.value})
	}


	render() {
	    return (
	    	<Query query={GET_NOT_TODO} notifyOnNetworkStatusChange>
		    	{({ loading, error, data, networkStatus }) => {

		    		if (networkStatus === 4) return "Refetching!";
		    		if (loading) return null;
		    		if (error) return `Error! ${error}`;

		    		const { getNotTodos = [] } = data
		    		const { name, description } = this.state
			    	return (
				      <div style={{display: 'flex', flexDirection: 'column', justifyContent: 'center', marginTop: '10%', width: '30%'}} >
				      <div style={{display: 'flex', flexDirection:'column'}}>
				        <h3 style={{margin: '0px'}}>Add something to not do!</h3>
				        <input placeholder='Name' name='name' onChange={this.onChange} value={name}/>
				        <input placeholder='Description' name='description' onChange={this.onChange} value={description}/>
						<Mutation mutation={CREATE_NOT_TODO} variables={{ name, description }} refetchQueries = {['GetNotTodos']}>
						  {createNotTodo => <button onClick={createNotTodo}>Submit</button>}
						</Mutation>				        
				      </div>
				      <table style={{border:'1px solid black', marginTop: '10%'}}>
				        <tbody>
				        <tr>
				          <th style={{border:'1px solid black'}}>name</th>
				          <th style={{border:'1px solid black'}}>description</th>
				          <th style={{border:'1px solid black'}}>complete</th>
				          <th style={{border:'1px solid black'}}>delete</th>
				        </tr>
				        {
				          getNotTodos.map(notTodo => (
				            <tr key={notTodo.id}>
				              <td style={{border:'1px solid black'}}>{notTodo.name}</td>
				              <td style={{border:'1px solid black'}}>{notTodo.description}</td>
				              <td style={{border:'1px solid black', textAlign: 'center'}}><input type="checkbox" key={notTodo.id} /></td>
				              <Mutation mutation={DELETE_TODO} variables={{ id: notTodo.id.substr(10,24) }} refetchQueries = {['GetNotTodos']}>
					              {deleteTodo => <td style={{border:'1px solid black', textAlign: 'center'}}><DeleteForeverIcon color="error" onClick={deleteTodo}/></td>}
							  </Mutation>	
				            </tr>
				          ))
				        }
				     	</tbody>
				       </table>
				     </div>
				    )
		    	}}
		    </Query>
		)    
  	}
}

export default Home