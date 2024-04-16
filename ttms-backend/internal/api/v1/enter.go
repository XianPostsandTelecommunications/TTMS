/**
 * @Author: lenovo
 * @Description:
 * @File:  enter
 * @Version: 1.0.0
 * @Date: 2023/05/29 8:42
 */

package v1

type group struct {
	User   user
	Email  email
	Movie  movie
	File   file
	Tag    tag
	Pm     pm
	Plan   plan
	Seat   seat
	Cinema cinema
	Ticket ticket
}

var Group = new(group)
